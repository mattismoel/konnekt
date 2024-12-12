import express from "express"
import routes from "@route/routes"
import { env } from "@/config/env";
import error from "./middleware/error";
import loggerMiddleware, { ConsoleLogger } from "./middleware/logger";
import cookieParser from "cookie-parser"
import cors from "cors"
import { S3ObjectStorage } from "./shared/object-storage/object-storage.s3";
import { createPermissionChecker } from "./middleware/rbac";
import { createEventController } from "./controller/event.controller";
import { createAuthController } from "./controller/auth.controller";
import { createGenreController } from "./controller/genre.controller";
import { createEventService } from "./service/event.service";
import { createAuthService } from "./service/auth.service";
import { createGenreService } from "./service/genre.service";
import { createRoleService } from "./service/role.service";
import { createSQLiteEventRepository } from "./repository/event.repository.sqlite";
import { createSQLiteSessionRepository } from "./repository/session.repository.sqlite";
import { createSQLiteUserRepository } from "./repository/user.repository.sqlite";
import { createSQLiteGenreRepository } from "./repository/genre.repository.sqlite";
import { createSQLiteRoleRepository } from "./repository/role.repository.sqlite";

const s3ObjectStorage = new S3ObjectStorage({
  bucket: "konnekt-bucket",
  region: "eu-north-1",
})

const eventRepo = createSQLiteEventRepository()
const eventService = createEventService(eventRepo, s3ObjectStorage)
const eventController = createEventController(eventService)

const sessionRepo = createSQLiteSessionRepository()
const userRepo = createSQLiteUserRepository()
const authService = createAuthService(sessionRepo, userRepo)

const authController = createAuthController(authService)

const genreRepo = createSQLiteGenreRepository()
const genreService = createGenreService(genreRepo)
const genreController = createGenreController(genreService)

const roleRepo = createSQLiteRoleRepository()
const roleService = createRoleService(roleRepo)

const logger = new ConsoleLogger()
const permissionChecker = createPermissionChecker(authService, roleService)

const app = express()

app.use(cors({ origin: env.FRONTEND_URL, credentials: true }))
app.use(loggerMiddleware(logger))
app.use(cookieParser())
app.use(express.json())

app.use(
  routes(
    eventController,
    authController,
    genreController,
    permissionChecker,
  )
);

app.use(error)

app.listen(env.PORT)
