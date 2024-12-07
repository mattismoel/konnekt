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
      console.error(e)
      next(e)
    }

    res.sendStatus(200)
  }

  update = async (req: Request, res: Response): Promise<void> => { }
  findAll = async (req: Request, res: Response): Promise<void> => { }
  findByID = async (req: Request, res: Response): Promise<void> => { }
}
