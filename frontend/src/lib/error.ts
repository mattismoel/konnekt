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
