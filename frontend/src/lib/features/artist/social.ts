import type { Component } from "svelte";

import MissingIcon from "~icons/mdi/question-mark-circle-outline"

import SpotifyIcon from "~icons/mdi/spotify"
import InstagramIcon from "~icons/mdi/instagram"
import AppleIcon from "~icons/mdi/apple"
import FacebookIcon from "~icons/mdi/facebook"
import YouTubeIcon from "~icons/mdi/youtube"

const iconMap = new Map<string, Component>([
	["spotify.com", SpotifyIcon],
	["instagram.com", InstagramIcon],
	["music.apple.com", AppleIcon],
	["facebook.com", FacebookIcon],
	["youtube.com", YouTubeIcon]
])

export const socialUrlToIcon = (url: string): Component => {
	const { hostname } = new URL(url);
	const iconKey = hostname.replace(/^www\./, '');

	const icon = iconMap.get(iconKey)

	if (!icon) return MissingIcon

	return icon
}
