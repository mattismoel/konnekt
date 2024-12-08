import type { NextFunction, Request, Response } from "express";
import type { AuthService } from "./auth.service";
import { AlreadyExistsError } from "@/shared/repo-error";
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
