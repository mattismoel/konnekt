import { Router } from "express";
import { type EventController } from "@/controller/event.controller";
import multer from "multer"
import type { PermissionCheckerMiddleware } from "@/middleware/rbac";

const upload = multer()

export const createEventRouter = (
  eventController: EventController,
  permissionChecker: PermissionCheckerMiddleware,
): Router => {
  const router = Router()

  router.get("/:id", eventController.getOneEvent)
  router.get("/", eventController.listEvents)
  router.post("/", permissionChecker(["event-create"]), eventController.createEvent)

  router.post("/coverImage",
    permissionChecker(["event-create", "event-update"]),
    upload.single("coverImage"),
    eventController.uploadCoverImage,
  )

  router.delete("/:id", permissionChecker(["event-delete"]), eventController.deleteEvent)

  return router
}
