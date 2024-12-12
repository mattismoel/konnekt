import type { RequestHandler } from "express";
import type { AuthService } from "@/service/auth.service";
import { AlreadyExistsError, NotFoundError } from "@/shared/repo-error";
import { SESSION_COOKIE_NAME } from "@/shared/auth/constant";
import { deleteSessionTokenCookie, setSessionTokenCookie } from "@/shared/auth/util";

export type AuthController = {
  register: RequestHandler
  logIn: RequestHandler
  logOut: RequestHandler
  validateSession: RequestHandler
}

export const createAuthController = (authService: AuthService): AuthController => {
  const register: RequestHandler = async (req, res, next) => {
    try {
      const { session, token, user } = await authService.register(req.body)
      setSessionTokenCookie(res, SESSION_COOKIE_NAME, token, session.expiresAt)
      res.send(user)
    } catch (e) {
      if (e instanceof AlreadyExistsError) {
        res.status(400).json({ error: "User already exists" })
        return
      }
      next(e)
    }
  }

  const logIn: RequestHandler = async (req, res, next) => {
    try {
      const prevToken = req.cookies[SESSION_COOKIE_NAME] as string;

      const { session: prevSession } = await authService.validateSessionToken(prevToken)
      if (prevSession) {
        await authService.invalidateSession(prevSession.id)
      }

      const { session, token, user } = await authService.login(req.body)
      setSessionTokenCookie(res, SESSION_COOKIE_NAME, token, session.expiresAt)
      res.send({ ...user, roles: user.roles.map(role => role.name) })
    } catch (e) {
      if (e instanceof NotFoundError) {
        res.status(401).json({ error: "User not found" })
        return
      }

      next(e)
    }
  }

  const logOut: RequestHandler = async (req, res) => {
    const token = req.cookies[SESSION_COOKIE_NAME] as string || null

    if (!token) {
      res.sendStatus(200)
      return
    }

    const { session } = await authService.validateSessionToken(token)

    if (!session) {
      res.sendStatus(200)
      return
    }

    await authService.invalidateSession(session.id)
    deleteSessionTokenCookie(res, SESSION_COOKIE_NAME)
    res.sendStatus(200)
  }

  const validateSession: RequestHandler = async (req, res) => {
    const token = req.cookies[SESSION_COOKIE_NAME] as string || null
    if (!token) {
      res.sendStatus(401)
      return
    }

    const { session, user } = await authService.validateSessionToken(token)

    if (!session) {
      res.sendStatus(401)
      return
    }

    res.send({ ...user, roles: user.roles.map(role => role.name) })
  }

  return {
    logIn,
    logOut,
    register,
    validateSession,
  }
}
