import { Router } from "express";
import { EventController } from "./event.controller";
import { SQLiteEventRepository } from "./event.repository.sqlite";
import { EventService } from "./event.service";

const eventRepository = new SQLiteEventRepository();
const eventService = new EventService(eventRepository)
const eventController = new EventController(eventService);

const router = Router()

router.get("/", async (req, res, next) => { })

router.post("/", eventController.create)
router.delete("/:id", eventController.delete)

export default router;
