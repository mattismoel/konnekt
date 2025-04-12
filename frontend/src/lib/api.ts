import { z, ZodSchema } from "zod"
import { APIError, apiErrorSchema } from "./error"

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

const querySchema = z.object({
	page: z
		.number()
		.int()
		.positive()
		.optional(),
	perPage: z
		.number()
		.int()
		.nonnegative()
		.optional(),
	limit: z
		.number()
		.int()
		.positive()
		.optional(),
	orderBy: z
		.map(z.string(), z.union([
			z.literal("ASC"),
			z.literal("DESC"),
		]))
		.optional(),
	filter: z
		.string()
		.array()
		.optional()
})

export type Query = z.infer<typeof querySchema>

export const createUrl = (base: string, query?: Query): URL => {
	return new URL(base + (query ? "?" + createQueryParams(query) : ""))
}

const createQueryParams = (query: Query): URLSearchParams => {
	const { page, perPage, orderBy, limit, filter } = querySchema.partial().parse(query)

	const params = new URLSearchParams()

	if (page) params.set("page", page.toString())

	if (perPage) params.set("per_page", perPage.toString())

	if (limit) params.set("limit", limit.toString())

	if (orderBy) {
		for (const [field, order] of orderBy?.entries()) {
			params.append("order_by", `${field} ${order}`)
		}
	}

	if (filter) {
		const filterClauses: string[] = []

		for (const filterStr of filter) {
			filterClauses.push(`${filterStr}`)
		}

		if (filterClauses.length > 0) {
			params.set("filter", filterClauses.join(","))
		}
	}

	return params
}
