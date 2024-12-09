import express from "express"
import routes from "@route/routes"
import { env } from "@/config/env";
import error from "./middleware/error";
import loggerMiddleware, { ConsoleLogger } from "./middleware/logger";
import cookieParser from "cookie-parser"
import cors from "cors"

const logger = new ConsoleLogger()

const app = express()

app.use(cors({ origin: env.FRONTEND_URL, credentials: true }))
app.use(loggerMiddleware(logger))
app.use(cookieParser())
app.use(express.json())

app.use(routes);
app.use(error)

app.listen(env.PORT)
