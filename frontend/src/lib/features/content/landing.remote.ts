import { getRequestEvent, query } from "$app/server";
import { landingImage } from "./landing";
import { createFileUrl, createListResult } from "$lib/pocketbase";
import { queryOptions } from "$lib/query";

export const getLandingImages = query(queryOptions.optional(), async (opts) => {
	const { locals } = getRequestEvent();

	const result = await locals.pb
		.collection("landingImages")
		.getList(opts?.page, opts?.perPage, { ...opts })
		.then((result) => createListResult(landingImage).parse(result))
		.then(({ items, ...result }) => ({
			...result,
			items: items.map((item) => ({
				...item,
				image: createFileUrl("landingImages", item.id, item.image)
			}))
		}));
	return result;
});
