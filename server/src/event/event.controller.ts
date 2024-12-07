import type { NextFunction, Request, Response } from "express";
import { ZodError } from "zod";
import type { EventService } from "./event.service";

export class EventController {
  constructor(
    private readonly eventService: EventService
  ) { }

  async create(req: Request, res: Response, next: NextFunction): Promise<void> {
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

  async delete(req: Request, res: Response, next: NextFunction): Promise<void> {
    try {
      const id = parseInt(req.params.id)
      await this.eventService.delete(id)
    } catch (e) {
      console.error(e)
      next(e)
    }

    res.sendStatus(200)
  }

  async update(req: Request, res: Response): Promise<void> { }
  async findAll(req: Request, res: Response): Promise<void> { }
  async findByID(req: Request, res: Response): Promise<void> { }
}
