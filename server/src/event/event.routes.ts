import { Router } from "express";
import { EventController } from "./event.controller";
import { SQLiteEventRepository } from "./event.repository.sqlite";
import { EventService } from "./event.service";
import multer from "multer"
import { S3ObjectStorage } from "@/shared/object-storage/object-storage.s3";
import { checkPermissions } from "@/middleware/rbac"
import { AuthService } from "@/auth/auth.service";
import { SQLiteSessionRepository } from "@/auth/session.repository.sqlite";
import { SQLiteUserRepository } from "@/auth/user.repository.sqlite";
import { RoleService } from "@/auth/role.service";
import { SQLiteRoleRepository } from "@/auth/role.repository.sqlite";


const s3ObjectStorage = new S3ObjectStorage({
  bucket: "konnekt-bucket",
  region: "eu-north-1",
})

const upload = multer()

const eventRepository = new SQLiteEventRepository();
const eventService = new EventService(eventRepository, s3ObjectStorage)
const eventController = new EventController(eventService);

const sessionRepository = new SQLiteSessionRepository()
const userRepository = new SQLiteUserRepository()
const roleRepository = new SQLiteRoleRepository()

const authService = new AuthService(sessionRepository, userRepository)
const roleService = new RoleService(roleRepository)

const router = Router()

router.get("/:id", eventController.getByID)
router.get("/", eventController.getAll)
router.post("/", checkPermissions(authService, roleService, ["event-create"]), eventController.create)

router.patch("/:id/coverImage",
  checkPermissions(authService, roleService, ["event-create", "event-update"]),
  upload.single("coverImage"),
  eventController.associateCoverImage,
)

router.delete("/:id", checkPermissions(authService, roleService, ["event-delete"]), eventController.delete)

export default router;
