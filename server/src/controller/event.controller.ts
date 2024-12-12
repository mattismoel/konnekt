import type { NextFunction, Request, RequestHandler, Response } from "express";
import { ZodError } from "zod";
import type { EventService } from "@/service/event.service";
import { NotFoundError } from "@/shared/repo-error";

export type EventController = {
  createEvent: RequestHandler;
  deleteEvent: RequestHandler;
  updateEvent: RequestHandler;
  listEvents: RequestHandler;
  getOneEvent: RequestHandler;
  uploadCoverImage: RequestHandler;
}

export const createEventController = (eventService: EventService): EventController => {
  const createEvent: RequestHandler = async (req, res, next): Promise<void> => {
    try {
      const event = await eventService.createEvent(req.body)
      res.json(event)
    } catch (e) {
      if (e instanceof ZodError) {
        res.status(400).json(e.flatten())
      } else {
        next(e)
      }
    }
  }

  const updateEvent: RequestHandler = async (req, res): Promise<void> => { }

  const deleteEvent: RequestHandler = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
      const id = parseInt(req.params.id)
      await eventService.deleteEvent(id)
    } catch (e) {
      next(e)
    }

    res.sendStatus(200)
  }

  const listEvents: RequestHandler = async (req, res, next): Promise<void> => {
    try {
      const events = await eventService.listEvents()
      res.json(events)
    } catch (e) {
      next(e)
    }
  }

  const getOneEvent: RequestHandler = async (req, res, next): Promise<void> => {
    try {
      const id = parseInt(req.params.id)
      const event = await eventService.getEventByID(id)
      res.json(event)
    } catch (e) {
      if (e instanceof NotFoundError) {
        res.sendStatus(404)
        return
      }

      next(e)
    }
  }

  const uploadCoverImage: RequestHandler = async (req, res, next): Promise<void> => {
    const file = req.file
    if (!file) {
      res.sendStatus(400)
      return
    }

    const url = await eventService.uploadCoverImage(file.buffer)
    res.setHeader("Content-Type", "text/plain")
    res.send(url)
  }

  return {
    createEvent,
    deleteEvent,
    listEvents,
    getOneEvent,
    uploadCoverImage,
    updateEvent
  }
}
