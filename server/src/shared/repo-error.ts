export class NotFoundError extends Error {
  constructor(msg: string) {
    super(`Not Found: ${msg}`)
    Object.setPrototypeOf(this, NotFoundError.prototype)
  }
}

export class AlreadyExistsError extends Error {
  constructor(msg: string) {
    super(`Already exists: ${msg}`)
    Object.setPrototypeOf(this, AlreadyExistsError.prototype)
  }
}
