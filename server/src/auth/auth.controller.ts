import type { NextFunction, Request, Response } from "express";
import type { AuthService } from "./auth.service";
import { AlreadyExistsError, NotFoundError } from "@/shared/repo-error";
import { SESSION_COOKIE_NAME } from "./constant";

export class AuthController {
  constructor(
    private readonly authService: AuthService
  ) { }


  register = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const { session, token } = await this.authService.register(req.body)
      this.setSessionTokenCookie(res, SESSION_COOKIE_NAME, token, session.expiresAt)
    } catch (e) {
      if (e instanceof AlreadyExistsError) {
        res.status(400).json({ error: "User already exists" })
        return
      }

      next(e)
    }

    res.sendStatus(201)
  }

  login = async (req: Request, res: Response, next: NextFunction) => {
    try {
      const prevToken = req.cookies[SESSION_COOKIE_NAME] as string;

      const { session: prevSession } = await this.authService.validateSessionToken(prevToken)

      if (prevSession) {
        this.setSessionTokenCookie(res, SESSION_COOKIE_NAME, prevToken, prevSession.expiresAt)
        res.sendStatus(200)
        return
      }

      const { session, token } = await this.authService.login(req.body)
      this.setSessionTokenCookie(res, SESSION_COOKIE_NAME, token, session.expiresAt)
      res.sendStatus(200)
    } catch (e) {
      if (e instanceof NotFoundError) {
        res.status(401).json({ error: "User not found" })
        return
      }

      next(e)
    }
  }

  logOut = async (req: Request, res: Response, next: NextFunction) => {
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
    this.deleteSessionTokenCookie(res, SESSION_COOKIE_NAME)
    res.sendStatus(200)
  }

  setSessionTokenCookie = (res: Response, name: string, token: string, expiresAt: Date) => {
    res.cookie(name, token, {
      httpOnly: true,
      sameSite: "lax",
      expires: expiresAt,
      path: "/"
    })
  }

  deleteSessionTokenCookie = (res: Response, name: string) => {
    res.cookie(name, "", {
      httpOnly: true,
      sameSite: "lax",
      maxAge: 0,
      path: "/"
    })
  }
}
