import { Router } from "express";
import { EventController } from "./event.controller";
import { SQLiteEventRepository } from "./event.repository.sqlite";
import { EventService } from "./event.service";

const eventRepository = new SQLiteEventRepository();
const eventService = new EventService(eventRepository)
const eventController = new EventController(eventService);

const router = Router()

router.post("/", async (req, res, next) => {
  console.log("Hello", req.body)
  await eventController.create(req, res, next)
})

export default router;
