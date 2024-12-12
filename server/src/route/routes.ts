import { Router } from "express";

import eventRoutes from "./event.routes"
import authRoutes from "./auth.routes"
import genreRoutes from "./genre.routes"
import type { GenreController } from "@/controller/genre.controller";
import type { AuthService } from "@/service/auth.service";
import type { RoleService } from "@/service/role.service";
import type { EventController } from "@/controller/event.controller";
import type { AuthController } from "@/controller/auth.controller";

const router = (
  eventController: EventController,
  authController: AuthController,
  genreController: GenreController,
  authService: AuthService,
  roleService: RoleService,
): Router => {
  const router = Router()
    .use("/events", eventRoutes(eventController, authService, roleService))
    .use("/auth", authRoutes(authController))
    .use("/genres", genreRoutes(genreController, authService, roleService))

  return router
}

export default router
