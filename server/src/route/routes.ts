import { Router } from "express";

import eventRoutes from "@/event/event.routes"

const router = Router()
  .use("/events", eventRoutes)

export default router;
