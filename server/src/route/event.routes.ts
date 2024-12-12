import { Router } from "express";
import { EventController } from "@/controller/event.controller";
import multer from "multer"
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";

const upload = multer()

const eventRouter = (
  eventController: EventController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {
  const router = Router()

  router.get("/:id", eventController.getByID)
  router.get("/", eventController.getAll)
  router.post("/", permissionChecker(["event-create"]), eventController.create)

  router.post("/coverImage",
    permissionChecker(["event-create", "event-update"]),
    upload.single("coverImage"),
    eventController.uploadCoverImage,
  )

  router.delete("/:id", permissionChecker(["event-delete"]), eventController.delete)

  return router
}

export default eventRouter;
