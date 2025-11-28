import { z } from "zod";
import { idSchema } from "./api";

export const auditDates = z.object({
  createdAt: z.coerce.date(),
  updatedAt: z.coerce.date(),
});

export const auditMembers = z.object({
  createdBy: idSchema,
  updatedBy: idSchema,
});
