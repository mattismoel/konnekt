import { Router } from "express";

import eventRoutes from "./event.routes"
import authRoutes from "./auth.routes"
import genreRoutes from "./genre.routes"
import type { GenreController } from "@/controller/genre.controller";
import type { EventController } from "@/controller/event.controller";
import type { AuthController } from "@/controller/auth.controller";
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";

const router = (
  eventController: EventController,
  authController: AuthController,
  genreController: GenreController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {
  const router = Router()
    .use("/events", eventRoutes(eventController, permissionChecker))
    .use("/auth", authRoutes(authController))
    .use("/genres", genreRoutes(genreController, permissionChecker))

  return router
}

export default router
