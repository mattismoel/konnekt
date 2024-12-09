import { checkPermissions } from "@/middleware/rbac";
import { Router } from "express";
import { GenreController } from "./genre.controller";
import { GenreService } from "./genre.service";
import { SQLiteGenreRepository } from "./genre.repository.sqlite";
import { AuthService } from "@/auth/auth.service";
import { RoleService } from "@/auth/role.service";
import { SQLiteSessionRepository } from "@/auth/session.repository.sqlite";
import { SQLiteUserRepository } from "@/auth/user.repository.sqlite";
import { SQLiteRoleRepository } from "@/auth/role.repository.sqlite";

const genreRepo = new SQLiteGenreRepository()
const sessionRepo = new SQLiteSessionRepository()
const userRepo = new SQLiteUserRepository()
const roleRepo = new SQLiteRoleRepository()


const genreService = new GenreService(genreRepo)
const controller = new GenreController(genreService)

const authService = new AuthService(sessionRepo, userRepo)
const roleService = new RoleService(roleRepo)

const router = Router()

router.get("/", checkPermissions(authService, roleService, ["genre-list"]), controller.getAll)

export default router;
