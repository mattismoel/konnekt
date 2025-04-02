import { z } from "zod"

export const apiErrorSchema = z.object({
	message: z.string(),
})

export class APIError extends Error {
	public readonly cause: string;
	public readonly status: number;

	constructor(status: number, message: string, cause: string) {
		super(message)
		this.name = "APIError"
		this.status = status
		this.cause = cause
		Object.setPrototypeOf(this, APIError.prototype)
	};
}

type Success<T> = { data: T, error: null }
type Failure<E> = { data: null, error: E }

type Result<T, E = Error> = Success<T> | Failure<E>

export const tryCatch = async <T, E = Error>(promise: Promise<T>): Promise<Result<T, E>> => {
	try {
		const data = await promise;
		return { data, error: null }
	} catch (e) {
		return { data: null, error: e as E }
	}
}
