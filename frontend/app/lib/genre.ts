import { z } from "zod"
import env from "@/config/env"

/**
 * @description Gets all genres and returns them as a string array.
 */
export const fetchAllGenres = async (): Promise<string[]> => {
  const res = await fetch(`${window.ENV.BACKEND_URL}/genres`, {
    credentials: "include"
  })

  if (!res.ok) {
    console.error(`Could not get genres: ${res.statusText}`)
    return []
  }

  const genres = z.string().array().parse(await res.json())

  return genres
}
