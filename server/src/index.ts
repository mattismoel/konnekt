import express from "express"
import { env } from "@/config/env";
import error from "./middleware/error";
const app = express()

app.use(express.json())
app.use(error)

app.listen(env.PORT)
