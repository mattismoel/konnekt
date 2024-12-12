import type { NextFunction, Request, RequestHandler, Response } from "express";
import { ZodError } from "zod";
import type { EventService } from "@/service/event.service";
import { NotFoundError } from "@/shared/repo-error";
import { DEFAULT_PAGE_SIZE } from "@/shared/event/constant";


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

  const listEvents: RequestHandler = async (req, res): Promise<void> => {
    const { page, limit, pageSize, search } = req.query

    //TODO: Implement maximum page size - fallback to DEFAULT_PAGE_SIZE.

    const events = await eventService.listEvents({
      page: page ? parseInt(page as string, 10) : 1,
      limit: limit ? parseInt(limit as string, 10) : undefined,
      pageSize: pageSize ? parseInt(pageSize as string, 10) : DEFAULT_PAGE_SIZE,
      search: search as string || undefined
    })

    res.json(events)
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
