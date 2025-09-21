import { FaGlobe } from "react-icons/fa6";
import type { IconType } from "react-icons/lib";
import { SiApple, SiFacebook, SiInstagram, SiSpotify, SiTidal, SiTiktok, SiYoutube } from "react-icons/si"

const iconMap = new Map<string, IconType>([
	["spotify", SiSpotify],
	["apple", SiApple],
	["tidal", SiTidal],
	["instagram", SiInstagram],
	["facebook", SiFacebook],
	["youtube", SiYoutube],
	["tiktok", SiTiktok],
])

export const socialIconByUrl = (url: string): IconType => {
	const [_, icon] = Array.from(iconMap).find(([name]) => url.includes(name)) ?? ["unknown", FaGlobe]

	return icon
}
