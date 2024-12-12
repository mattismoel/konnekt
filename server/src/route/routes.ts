import { Router } from "express";

import type { GenreController } from "@/controller/genre.controller";
import type { EventController } from "@/controller/event.controller";
import type { AuthController } from "@/controller/auth.controller";
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";
import { createEventRouter } from "./event.routes";
import { createAuthRouter } from "./auth.routes";
import { createGenreRouter } from "./genre.routes";

const router = (
  eventController: EventController,
  authController: AuthController,
  genreController: GenreController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {
  const router = Router()
    .use("/events", createEventRouter(eventController, permissionChecker))
    .use("/auth", createAuthRouter(authController))
    .use("/genres", createGenreRouter(genreController, permissionChecker))

  return router
}

export default router
