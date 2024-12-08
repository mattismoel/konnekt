import { sha256 } from "@oslojs/crypto/sha2";
import { registerSchema, type RegisterDTO } from "./auth.dto";
import type { Session, SessionValidationResult } from "./session.dto";
import type { SessionRepository } from "./session.repository";
import type { UserRepository } from "./user.repository";
import { encodeBase32NoPadding, encodeHexLowerCase } from "@oslojs/encoding"
import { addDays, isAfter, subDays } from "date-fns";
import { SESSION_COOKIE_NAME, SESSION_LIFETIME_DAYS, SESSION_REFRESH_DAYS } from "./constant";
import type { Response } from "express";
import { hash } from "@node-rs/argon2";
import { AlreadyExistsError } from "@/shared/repo-error";

export class AuthService {
  constructor(
    private readonly sessionRepository: SessionRepository,
    private readonly userRepository: UserRepository,
  ) { }

  register = async (res: Response, data: RegisterDTO) => {
    const { password, passwordConfirm, ...userData } = registerSchema.parse(data)

    const existingUser = await this.userRepository.getByEmail(userData.email)
    if (existingUser) {
      throw new AlreadyExistsError(`User with email ${userData.email}`)
    }

    const passwordHash = await hash(password, {
      memoryCost: 19456,
      timeCost: 2,
      outputLen: 32,
      parallelism: 1
    })

    const user = await this.userRepository.insert({ ...userData, passwordHash })

    const token = this.generateSessionToken()
    const session = await this.createSession(token, user.id)

    this.setSessionTokenCookie(res, token, session.expiresAt)
  }

  generateSessionToken = (): string => {
    const bytes = new Uint8Array(20);
    crypto.getRandomValues(bytes);
    const token = encodeBase32NoPadding(bytes)
    return token
  }

  createSession = async (token: string, userID: number): Promise<Session> => {
    const sessionID = this.sessionIDFromToken(token)

    const session: Session = {
      id: sessionID,
      userID,
      expiresAt: addDays(new Date(), SESSION_LIFETIME_DAYS)
    }

    await this.sessionRepository.insert(session)
    return session
  }

  validateSessionToken = async (token: string): Promise<SessionValidationResult> => {
    const sessionID = this.sessionIDFromToken(token)
    const session = await this.sessionRepository.getByID(sessionID)

    if (!session) {
      return { session: null, user: null }
    }

    const user = await this.userRepository.getByID(session.userID)

    if (!user) {
      return { session: null, user: null }
    }

    const now = new Date()

    if (isAfter(now, session.expiresAt)) {
      await this.sessionRepository.delete(session.id)
      return { session: null, user: null }
    }

    if (isAfter(now, subDays(session.expiresAt, SESSION_REFRESH_DAYS))) {
      session.expiresAt = addDays(now, SESSION_LIFETIME_DAYS)
      this.sessionRepository.setExpiry(session.id, session.expiresAt)
      return { session, user }
    }

    return { session, user }
  }

  invalidateSession = async (sessionID: string) => {
    await this.sessionRepository.delete(sessionID)
  }

  sessionIDFromToken = (token: string) => {
    const sessionID = encodeHexLowerCase(sha256(new TextEncoder().encode(token)))
    return sessionID
  }

  setSessionTokenCookie = (res: Response, token: string, expiresAt: Date) => {
    res.cookie(SESSION_COOKIE_NAME, token, {
      httpOnly: true,
      sameSite: "lax",
      expires: expiresAt,
      path: "/"
    })
  }

  deleteSessionTokenCookie = (res: Response) => {
    res.cookie(SESSION_COOKIE_NAME, "", {
      httpOnly: true,
      sameSite: "lax",
      maxAge: 0,
      path: "/"
    })
  }
}
