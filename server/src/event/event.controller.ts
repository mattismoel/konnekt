import type { NextFunction, Request, Response } from "express";
import { ZodError } from "zod";
import type { EventService } from "./event.service";
import { NotFoundError } from "./event.repository";

export class EventController {
  constructor(
    private readonly eventService: EventService
  ) { }

  create = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
      const event = await this.eventService.create(req.body)
      res.json(event)
    } catch (e) {
      if (e instanceof ZodError) {
        res.status(400).json(e.flatten())
      } else {
        next(e)
      }
    }
  }

  delete = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
      const id = parseInt(req.params.id)
      await this.eventService.delete(id)
    } catch (e) {
      next(e)
    }

    res.sendStatus(200)
  }

  getAll = async (_req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
      const events = await this.eventService.getAll()
      res.json(events)
    } catch (e) {
      next(e)
    }
  }

  getByID = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    try {
      const id = parseInt(req.params.id)
      const event = await this.eventService.getByID(id)
      res.json(event)
    } catch (e) {
      if (e instanceof NotFoundError) {
        res.sendStatus(404)
        return
      }

      next(e)
    }
  }

  update = async (req: Request, res: Response): Promise<void> => { }

  associateCoverImage = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
    const id = parseInt(req.params.id)
    const file = req.file
    if (!file) {
      res.sendStatus(400)
      return
    }

    await this.eventService.setCoverImage(id, file.buffer)
    res.sendStatus(200)
  }
}
