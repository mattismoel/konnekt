import express from "express"
import routes from "@route/routes"
import { env } from "@/config/env";
import error from "./middleware/error";
import loggerMiddleware, { ConsoleLogger } from "./middleware/logger";
import cookieParser from "cookie-parser"
import cors from "cors"
import { S3ObjectStorage } from "./shared/object-storage/object-storage.s3";
import { EventController } from "./controller/event.controller";
import { AuthController } from "./controller/auth.controller";
import { EventService } from "./service/event.service";
import { SQLiteEventRepository } from "./repository/event.repository.sqlite";
import { AuthService } from "./service/auth.service";
import { SQLiteSessionRepository } from "./repository/session.repository.sqlite";
import { SQLiteUserRepository } from "./repository/user.repository.sqlite";
import { RoleService } from "./service/role.service";
import { GenreController } from "./controller/genre.controller";
import { GenreService } from "./service/genre.service";
import { SQLiteGenreRepository } from "./repository/genre.repository.sqlite";
import { SQLiteRoleRepository } from "./repository/role.repository.sqlite";

const s3ObjectStorage = new S3ObjectStorage({
  bucket: "konnekt-bucket",
  region: "eu-north-1",
})

const eventRepo = new SQLiteEventRepository()
const eventService = new EventService(eventRepo, s3ObjectStorage)
const eventController = new EventController(eventService)

const sessionRepo = new SQLiteSessionRepository()
const userRepo = new SQLiteUserRepository()
const authService = new AuthService(sessionRepo, userRepo)

const authController = new AuthController(authService)

const genreRepo = new SQLiteGenreRepository()
const genreService = new GenreService(genreRepo)
const genreController = new GenreController(genreService)

const roleRepo = new SQLiteRoleRepository()
const roleService = new RoleService(roleRepo)

const logger = new ConsoleLogger()

const app = express()

app.use(cors({ origin: env.FRONTEND_URL, credentials: true }))
app.use(loggerMiddleware(logger))
app.use(cookieParser())
app.use(express.json())

app.use(routes(eventController, authController, genreController, authService, roleService));
app.use(error)

app.listen(env.PORT)
