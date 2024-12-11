import { Router } from "express";
import { AuthController } from "@/controller/auth.controller";
import { AuthService } from "@/service/auth.service";
import { SQLiteSessionRepository } from "@/repository/session.repository.sqlite";
import { SQLiteUserRepository } from "@/repository/user.repository.sqlite";
import { UserService } from "@/service/user.service";

const sessionRepository = new SQLiteSessionRepository()
const userRepository = new SQLiteUserRepository()

const userService = new UserService(userRepository)
const authService = new AuthService(sessionRepository, userRepository)
const authController = new AuthController(authService, userService)

const router = Router()

router.post("/register", authController.register)
router.post("/login", authController.login)
router.post("/log-out", authController.logOut)
router.get("/validate-session", authController.validateSession)

export default router;
