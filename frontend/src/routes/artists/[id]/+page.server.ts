import { artistById } from "$lib/artist";
import type { PageServerLoad } from "./$types";

export const load: PageServerLoad = async ({ params }) => {
  const artist = await artistById(parseInt(params.id))

  return { artist }
}
