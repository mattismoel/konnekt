import type { WithClassName } from "../class"
import { cn } from "../clsx"
import SocialIcon from "./social-icon"

type Props = WithClassName<{
	urls: string[]
	size?: "sm" | "md" | "lg"
}>

const SocialList = ({ urls, size = "md", className }: Props) => (
	<ul className={cn("flex items-center", {
		"text-xl gap-4": size === "sm",
		"text-2xl gap-5": size === "md",
		"text-3xl gap-6": size === "lg",
	}, className)}>
		{urls.map(url => <li><SocialIcon url={url} /></li>)}
	</ul>
)

export default SocialList
