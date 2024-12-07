import express from "express"
import routes from "@route/routes"
import { env } from "@/config/env";
import error from "./middleware/error";

const app = express()

app.use(express.json())

app.use(routes);
app.use(error)

app.listen(env.PORT)
