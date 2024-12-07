import type { NextFunction, Request, Response } from "express";
import { ZodError } from "zod";
import type { EventService } from "./event.service";

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
        res.json(e.flatten())
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

  getAll = async (req: Request, res: Response, next: NextFunction): Promise<void> => {
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
      next(e)
    }
  }

  update = async (req: Request, res: Response): Promise<void> => { }
}
