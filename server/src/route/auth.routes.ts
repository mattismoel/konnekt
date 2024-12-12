import { Router } from "express";
import { AuthController } from "@/controller/auth.controller";

const authRouter = (authController: AuthController): Router => {
  const router = Router()

  router.post("/register", authController.register)
  router.post("/login", authController.login)
  router.post("/log-out", authController.logOut)
  router.get("/validate-session", authController.validateSession)

  return router
}

export default authRouter;
