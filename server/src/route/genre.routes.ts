import { checkPermissions } from "@/middleware/rbac";
import { Router } from "express";
import { GenreController } from "@/controller/genre.controller";
import { GenreService } from "@/service/genre.service";
import { SQLiteGenreRepository } from "@/repository/genre.repository.sqlite";
import { AuthService } from "@/service/auth.service";
import { RoleService } from "@/service/role.service";
import { SQLiteSessionRepository } from "@/repository/session.repository.sqlite";
import { SQLiteUserRepository } from "@/repository/user.repository.sqlite";
import { SQLiteRoleRepository } from "@/repository/role.repository.sqlite";

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
