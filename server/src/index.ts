import express from "express"
import { env } from "@/config/env";
const app = express()

app.use(express.json())
app.listen(env.PORT)
