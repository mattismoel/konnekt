import { Router } from "express";

import eventRoutes from "@/event/event.routes"
import authRoutes from "@/auth/auth.routes"
import genreRoutes from "@/genre/genre.routes"

const router = Router()
  .use("/events", eventRoutes)
  .use("/auth", authRoutes)
  .use("/genres", genreRoutes)

export default router;
