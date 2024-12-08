import { Router } from "express";

import eventRoutes from "@/event/event.routes"
import authRoutes from "@/auth/auth.routes"

const router = Router()
  .use("/events", eventRoutes)
  .use("/auth", authRoutes)

export default router;
