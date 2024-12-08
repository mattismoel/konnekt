import { Router } from "express";
import { AuthController } from "./auth.controller";
import { AuthService } from "./auth.service";
import { SQLiteSessionRepository } from "./session.repository.sqlite";
import { SQLiteUserRepository } from "./user.repository.sqlite";

const sessionRepository = new SQLiteSessionRepository()
const userRepository = new SQLiteUserRepository()

const authService = new AuthService(sessionRepository, userRepository)
const authController = new AuthController(authService)

const router = Router()

router.post("/register", authController.register)
router.post("/login", authController.login)

export default router;
