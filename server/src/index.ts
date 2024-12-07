import express from "express"
import routes from "@route/routes"
import { env } from "@/config/env";
import error from "./middleware/error";
import loggerMiddleware, { ConsoleLogger } from "./middleware/logger";

const logger = new ConsoleLogger()

const app = express()

app.use(loggerMiddleware(logger))
app.use(express.json())

app.use(routes);
app.use(error)

app.listen(env.PORT)
