import { Router } from "express";
import { type GenreController } from "@/controller/genre.controller";
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";

export const createGenreRouter = (
  controller: GenreController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {

  const router = Router()
  router.get("/", permissionChecker(["genre-list"]), controller.listGenres)

  return router
}
