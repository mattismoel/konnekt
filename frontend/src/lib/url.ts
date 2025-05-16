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

export const createUrl = (base: string, query?: Query): string => {
  // let origin: string = ""
  //
  // When we are server side, the nginx container will be directly acessible
  // over the shared Docker network by its 'nginx' container name.
  // if (!browser) {
  // console.log("NOT BROWSER")
  // origin = "http://backend:8080"
  // }

  const url = base + (query ? "?" + createQueryParams(query) : "")

  return url
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

/**
 * @description Cleans the input URL for any trailing slashes
 */
export const cleanUrl = (url: string): string => {
  return url.replace(/\/$/, "")
}

/**
 * @description Checks if an input string is a valid URL.
 */
export const isValidUrl = (s: string): boolean => {
  let url: URL;
  try {
    url = new URL(s)
  } catch (_) {
    return false
  }

  return url.protocol === "http:" || url.protocol === "https:"
}
