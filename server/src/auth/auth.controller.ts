import type { NextFunction, Request, Response } from "express";
import type { AuthService } from "./auth.service";
import { AlreadyExistsError } from "@/shared/repo-error";

export class AuthController {
  constructor(
    private readonly authService: AuthService
  ) { }


  register = async (req: Request, res: Response, next: NextFunction) => {
    try {
      await this.authService.register(res, req.body)
    } catch (e) {
      if (e instanceof AlreadyExistsError) {
        res.status(400).json({ error: "User already exists" })
        return
      }

      next(e)
    }

    res.sendStatus(201)
  }
}
