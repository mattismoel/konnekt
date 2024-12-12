import { Router } from "express";
import { EventController } from "@/controller/event.controller";
import multer from "multer"
import { checkPermissions } from "@/middleware/rbac"
import { AuthService } from "@/service/auth.service";
import { RoleService } from "@/service/role.service";

const upload = multer()

const eventRouter = (eventController: EventController, authService: AuthService, roleService: RoleService): Router => {
  const router = Router()

  router.get("/:id", eventController.getByID)
  router.get("/", eventController.getAll)
  router.post("/", checkPermissions(authService, roleService, ["event-create"]), eventController.create)

  router.post("/coverImage",
    checkPermissions(authService, roleService, ["event-create", "event-update"]),
    upload.single("coverImage"),
    eventController.uploadCoverImage,
  )

  router.delete("/:id", checkPermissions(authService, roleService, ["event-delete"]), eventController.delete)

  return router
}

export default eventRouter;
