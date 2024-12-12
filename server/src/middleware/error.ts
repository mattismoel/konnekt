import type { ErrorRequestHandler } from "express";

const error: ErrorRequestHandler = (err, req, res, next) => {
  console.error(err);
  res.status(500).send("Something went wrong.")
}

export default error;
