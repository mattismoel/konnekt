import { isValidUrl } from "./url"

/**
 * @description Retrieves the track ID from a standard Spotify track URL (from 
 * the "Share" menu).
 * @param {string} url - The standard sharable URL from the Spotify "Share"
 * menu
 * @returns {(string|null)} Track ID of Spotify track.
 */
export const spotifyTrackIdFromUrl = (url: string): (string | null) => {
	if (!isValidUrl(url)) return null

	const { pathname } = new URL(url)
	const id = pathname.split("/").at(2)

	return id ?? null
}
