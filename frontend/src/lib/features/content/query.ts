import { queryOptions } from "@tanstack/react-query";
import { landingImages } from "./landing";

export const landingImagesQueryOptions = queryOptions({
	queryKey: ["landing", "images"],
	queryFn: landingImages
})
