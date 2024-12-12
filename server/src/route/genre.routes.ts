import { Router } from "express";
import { GenreController } from "@/controller/genre.controller";
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";

const genreRouter = (
  controller: GenreController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {

  const router = Router()
  router.get("/", permissionChecker(["genre-list"]), controller.getAll)

  return router
}

export default genreRouter;
