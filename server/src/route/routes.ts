import { Router } from "express";

import eventRoutes from "./event.routes"
import authRoutes from "./auth.routes"
import genreRoutes from "./genre.routes"

const router = Router()
  .use("/events", eventRoutes)
  .use("/auth", authRoutes)
  .use("/genres", genreRoutes)

export default router;
