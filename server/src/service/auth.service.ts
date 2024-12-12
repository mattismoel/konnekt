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

export type AuthService = {
  register(data: RegisterDTO): Promise<{ session: Session, token: string, user: UserDTO }>;
  login(data: LoginDTO): Promise<{ session: Session, token: string, user: UserDTO }>
  validateSessionToken(token: string): Promise<SessionValidationResult>
  invalidateSession(sessionID: string): Promise<void>
}

export const createAuthService = (
  sessionRepo: SessionRepository,
  userRepo: UserRepository,
): AuthService => {
  const register = async (data: RegisterDTO): Promise<
    { session: Session, token: string, user: UserDTO }
  > => {
    const { password, passwordConfirm, ...userData } = registerSchema.parse(data)

    const existingUser = await userRepo.getUserByEmail(userData.email)
    if (existingUser) {
      throw new AlreadyExistsError(`User with email ${userData.email}`)
    }

    const passwordHash = await hash(password, {
      memoryCost: 19456,
      timeCost: 2,
      outputLen: 32,
      parallelism: 1
    })

    const user = await userRepo.insertUser({
      ...userData,
      passwordHash,
      roles: ["user"]
    })

    const token = generateSessionToken()
    const session = await createSession(token, user.id)

    await sessionRepo.insertSession(session)

    return { session, token, user }
  }

  const login = async (data: LoginDTO): Promise<{ session: Session, token: string, user: UserDTO }> => {
    const { email, password } = loginSchema.parse(data)

    const user = await userRepo.getUserByEmail(email)
    if (!user) {
      throw new NotFoundError(`User with email ${email}`)
    }

    const userPasswordHash = await userRepo.getUserPasswordHash(user.id)

    if (!userPasswordHash) {
      throw new NotFoundError(`Password hash for ${email}`)
    }

    if (!verify(userPasswordHash, password)) {
      throw new Error("Invalid password")
    }

    const token = generateSessionToken()
    const session = await createSession(token, user.id)

    await sessionRepo.insertSession(session)

    return { session, token, user }
  }
  const validateSessionToken = async (token: string): Promise<SessionValidationResult> => {
    const sessionID = sessionIDFromToken(token)
    const session = await sessionRepo.getSessionByID(sessionID)

    if (!session) {
      return { session: null, user: null }
    }

    const user = await userRepo.getUserByID(session.userID)

    if (!user) {
      return { session: null, user: null }
    }

    const now = new Date()

    if (isAfter(now, session.expiresAt)) {
      await sessionRepo.deleteSession(session.id)
      return { session: null, user: null }
    }

    if (isAfter(now, subDays(session.expiresAt, SESSION_REFRESH_DAYS))) {
      session.expiresAt = addDays(now, SESSION_LIFETIME_DAYS)
      await sessionRepo.setSessionExpiry(session.id, session.expiresAt)
      return { session, user }
    }

    return { session, user }
  }

  const invalidateSession = async (sessionID: string) => {
    await sessionRepo.deleteSession(sessionID)
  }

  return {
    register,
    login,
    validateSessionToken,
    invalidateSession,
  }
}
