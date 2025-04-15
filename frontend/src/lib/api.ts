import { z, ZodSchema } from "zod"

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

type Method = "GET" | "POST" | "PUT" | "PATCH" | "DELETE"

type RequestBody<T> = {
	bodySchema: Zod.Schema<T>
	body: T,
}

// requestAndParse("api.example.com", )

export async function requestAndParse<TBody, TResponse>(
	fetchFn: typeof fetch,
	url: string | URL,
	resSchema: ZodSchema<TResponse>,
	errorMsg?: string,
	bodyOpts?: RequestBody<TBody>,
	method?: Method
): Promise<TResponse>

export async function requestAndParse<TBody>(
	fetchFn: typeof fetch,
	url: string | URL,
	resSchema: undefined,
	errorMsg?: string,
	bodyOpts?: RequestBody<TBody>,
	method?: Method
): Promise<void>;

export async function requestAndParse<TBody, TResponse>(
	fetchFn: typeof fetch,
	url: string | URL,
	resSchema?: Zod.Schema<TResponse>,
	errorMsg?: string,
	bodyOpts?: RequestBody<TBody>,
	method?: Method,
): Promise<TResponse | void> {
	const res = await fetchFn(url, {
		headers: {
			"Content-Type": "application/json",
		},
		method,
		body: bodyOpts
			? JSON.stringify(bodyOpts.bodySchema.parse(bodyOpts.bodySchema))
			: undefined
	})

	if (!res.ok) {
		const err = apiErrorSchema.parse(await res.json())

		throw new APIError(
			res.status,
			errorMsg || "Something went wrong...",
			err.message,
		)
	}

	if (!resSchema) return

	const data = resSchema.parse(await res.json())
	return data
}


