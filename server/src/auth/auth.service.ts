import { sha256 } from "@oslojs/crypto/sha2";
import { loginSchema, registerSchema, type LoginDTO, type RegisterDTO } from "./auth.dto";
import type { Session, SessionValidationResult } from "./session.dto";
import type { SessionRepository } from "./session.repository";
import type { UserRepository } from "./user.repository";
import { encodeBase32NoPadding, encodeHexLowerCase } from "@oslojs/encoding"
import { addDays, isAfter, subDays } from "date-fns";
import { SESSION_LIFETIME_DAYS, SESSION_REFRESH_DAYS } from "./constant";
import { hash, verify } from "@node-rs/argon2";
import { AlreadyExistsError, NotFoundError } from "@/shared/repo-error";

export class AuthService {
  constructor(
    private readonly sessionRepository: SessionRepository,
    private readonly userRepository: UserRepository,
  ) { }

  /**
  * @description Attempts to register a user, returning the users newly created session and token.
  */
  register = async (data: RegisterDTO): Promise<
    { session: Session, token: string }
  > => {
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

    return { session, token }
  }

  login = async (data: LoginDTO): Promise<{ session: Session, token: string }> => {
    const { email, password } = loginSchema.parse(data)

    const user = await this.userRepository.getByEmail(email)
    if (!user) {
      throw new NotFoundError(`User with email ${email}`)
    }

    const userPasswordHash = await this.userRepository.getPasswordHash(user.id)

    if (!verify(userPasswordHash, password)) {
      throw new Error("Invalid password")
    }

    const token = this.generateSessionToken()
    const session = await this.createSession(token, user.id)

    return { session, token }
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
}
