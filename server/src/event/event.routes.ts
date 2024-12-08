import { Router } from "express";
import { EventController } from "./event.controller";
import { SQLiteEventRepository } from "./event.repository.sqlite";
import { EventService } from "./event.service";
import multer from "multer"
import { S3ObjectStorage } from "@/shared/object-storage/object-storage.s3";


const s3ObjectStorage = new S3ObjectStorage({
  bucket: "konnekt-bucket",
  region: "eu-north-1",
})

const upload = multer()

const eventRepository = new SQLiteEventRepository();
const eventService = new EventService(eventRepository, s3ObjectStorage)
const eventController = new EventController(eventService);

const router = Router()

router.get("/:id", eventController.getByID)
router.get("/", eventController.getAll)
router.post("/", eventController.create)
router.patch("/:id/coverImage", upload.single("coverImage"), eventController.associateCoverImage)
router.delete("/:id", eventController.delete)

export default router;
