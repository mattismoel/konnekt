import z from "zod";

const apiError = z.object({
	status: z.int().positive(),
	message: z.string().nonempty(),
	cause: z.string().nonempty(),
	path: z.string().nonempty(),
})

export const throwAPIError = (errData: any) => {
	const err = apiError.parse(errData)
	throw new Error(`${err.status}: ${err.message}: ${err.cause}`)
}
