import { isValidUrl } from "./url"

/**
 * @description Retrieves the track ID from a standard Spotify track URL (from 
 * the "Share" menu).
 * @param {string} urlStr - The standard sharable URL from the Spotify "Share"
 * menu
 * @returns {(string|null)} Track ID of Spotify track.
 */
export const trackIdFromUrl = (urlStr: string): string | null => {
  // https://open.spotify.com/track/{trackId}?si=54e63e520dfb4351
  if (!isValidUrl(urlStr)) return null

  const url = new URL(urlStr)

  // /track/{trackId} -> {trackId}
  const id = url.pathname.split("/").at(2)

  return id || null
}
