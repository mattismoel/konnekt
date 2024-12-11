import type { NextFunction, Request, RequestHandler, Response } from "express";
import type { AuthService } from "@/service/auth.service";
import { AlreadyExistsError, NotFoundError } from "@/shared/repo-error";
import { SESSION_COOKIE_NAME } from "@/shared/auth/constant";
import { deleteSessionTokenCookie, setSessionTokenCookie } from "@/shared/auth/util";

export class AuthController {
  constructor(
    private readonly authService: AuthService,
  ) { }

  register = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { session, token, user } = await this.authService.register(req.body)
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

  login = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const prevToken = req.cookies[SESSION_COOKIE_NAME] as string;

      const { session: prevSession } = await this.authService.validateSessionToken(prevToken)
      if (prevSession) {
        this.authService.invalidateSession(prevSession.id)
      }

      const { session, token, user } = await this.authService.login(req.body)
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

  logOut = async (req: Request, res: Response) => {
    const token = req.cookies[SESSION_COOKIE_NAME] as string || null

    if (!token) {
      res.sendStatus(200)
      return
    }

    const { session } = await this.authService.validateSessionToken(token)

    if (!session) {
      res.sendStatus(200)
      return
    }

    await this.authService.invalidateSession(session.id)
    deleteSessionTokenCookie(res, SESSION_COOKIE_NAME)
    res.sendStatus(200)
  }

  validateSession: RequestHandler = async (req, res) => {
    const token = req.cookies[SESSION_COOKIE_NAME] as string || null
    if (!token) {
      res.sendStatus(401)
      return
    }

    const { session, user } = await this.authService.validateSessionToken(token)

    if (!session) {
      res.sendStatus(401)
      return
    }

    res.send({ ...user, roles: user.roles.map(role => role.name) })
  }
}
