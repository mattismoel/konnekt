import { checkPermissions } from "@/middleware/rbac";
import { Router } from "express";
import { GenreController } from "@/controller/genre.controller";
import { AuthService } from "@/service/auth.service";
import { RoleService } from "@/service/role.service";

const genreRouter = (controller: GenreController, authService: AuthService, roleService: RoleService): Router => {

  const router = Router()
  router.get("/", checkPermissions(authService, roleService, ["genre-list"]), controller.getAll)

  return router
}

export default genreRouter;
