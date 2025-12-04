import { form } from "$app/server";
import z from "zod";

const testForm = z.object({
	fileField: z.file().max(5_000_000)
});

export const doTest = form(testForm, async (data) => {
	console.log(data);
});
