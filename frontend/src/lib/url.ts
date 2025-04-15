import { z } from "zod"

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
