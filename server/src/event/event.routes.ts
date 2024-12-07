import { Router } from "express";
import { EventController } from "./event.controller";
import { SQLiteEventRepository } from "./event.repository.sqlite";
import { EventService } from "./event.service";

const eventRepository = new SQLiteEventRepository();
const eventService = new EventService(eventRepository)
const eventController = new EventController(eventService);

const router = Router()

router.post("/", async (req, res, next) => {
  await eventController.create(req, res, next)
})

router.delete("/:id", async (req, res, next) => {
  await eventController.delete(req, res, next)
  res.sendStatus(200)
})

export default router;
