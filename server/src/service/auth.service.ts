import { loginSchema, registerSchema, type LoginDTO, type RegisterDTO } from "@/dto/auth.dto";
import type { Session, SessionValidationResult } from "@/dto/session.dto";
import type { SessionRepository } from "@/repository/session.repository";
import type { UserRepository } from "@/repository/user.repository";
import { addDays, isAfter, subDays } from "date-fns";
import { hash, verify } from "@node-rs/argon2";
import { AlreadyExistsError, NotFoundError } from "@/shared/repo-error";
import { SESSION_LIFETIME_DAYS, SESSION_REFRESH_DAYS } from "@/shared/auth/constant";
import { createSession, generateSessionToken, sessionIDFromToken } from "@/shared/auth/util";
import type { UserDTO } from "@/dto/user.dto";

export class AuthService {
  constructor(
    private readonly sessionRepository: SessionRepository,
    private readonly userRepository: UserRepository,
  ) { }

  /**
  * @description Attempts to register a user, returning the users newly created session and token.
  */
  register = async (data: RegisterDTO): Promise<
    { session: Session, token: string, user: UserDTO }
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

    const user = await this.userRepository.insert({ ...userData, passwordHash, roles: ["user"] })

    const token = generateSessionToken()
    const session = await createSession(token, user.id)

    await this.sessionRepository.insert(session)

    return { session, token, user }
  }

  /**
  * @description Attempts to log out the user. 
  * If successful, the created session and token is returned.
  */
  login = async (data: LoginDTO): Promise<{ session: Session, token: string, user: UserDTO }> => {
    const { email, password } = loginSchema.parse(data)

    const user = await this.userRepository.getByEmail(email)
    if (!user) {
      throw new NotFoundError(`User with email ${email}`)
    }

    const userPasswordHash = await this.userRepository.getPasswordHash(user.id)

    if (!userPasswordHash) {
      throw new NotFoundError(`Password hash for ${email}`)
    }

    if (!verify(userPasswordHash, password)) {
      throw new Error("Invalid password")
    }

    const token = generateSessionToken()
    const session = await createSession(token, user.id)

    await this.sessionRepository.insert(session)

    return { session, token, user }
  }

  /**
  * @description Validates an input session token. 
  *
  * - If the session is non-existent a null-like response is returned.
  * - If the session is to be refreshed, the the refreshed session is returned.
  * - If the session has expired, the session is invalidated and a null-like response is returned.
  */
  validateSessionToken = async (token: string): Promise<SessionValidationResult> => {
    const sessionID = sessionIDFromToken(token)
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

  /**
  * @description Invalidates the session, deleting it from the session repository
  */
  invalidateSession = async (sessionID: string) => {
    await this.sessionRepository.delete(sessionID)
  }
}
