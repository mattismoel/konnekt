import type { AuthController } from "@/controller/auth.controller";
import { Router } from "express";

export const createAuthRouter = (authController: AuthController): Router => {
  const router = Router()

  router.post("/register", authController.register)
  router.post("/login", authController.logIn)
  router.post("/log-out", authController.logOut)
  router.get("/validate-session", authController.validateSession)

  return router
}
