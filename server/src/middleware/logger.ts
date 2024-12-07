import { differenceInMilliseconds } from "date-fns";
import { type RequestHandler } from "express";

interface Logger {
  log(message: string): void
  error(e: Error): void
}

const loggerMiddleware = (logger: Logger): RequestHandler => {
  return (req, res, next) => {
    const startTime = new Date()

    const { method, path } = req

    const reqMsg = `<- ${method} ${path} | ${new Date().toISOString()}`
    logger.log(reqMsg)

    res.on("finish", () => {
      const durationMs = differenceInMilliseconds(new Date(), startTime)
      const resMsg = `-> ${method} ${path} ${res.statusCode} ${durationMs} ms`
      logger.log(resMsg)
    })

    next();
  }
}

export class ConsoleLogger implements Logger {
  log(message: string) {
    console.log(message)
  }

  error(e: Error): void {
    console.error(e)
  }
}

export default loggerMiddleware
