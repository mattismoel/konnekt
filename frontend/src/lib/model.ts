import z from "zod";

export const id = z.string().nonempty();

export const baseFields = z.object({ id: id, created: z.coerce.date(), updated: z.coerce.date() });
