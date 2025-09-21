import type { WithClassName } from "../class"
import { cn } from "../clsx"
import { socialIconByUrl } from "../social"

type Props = WithClassName<{
	url: string
}>

const SocialIcon = ({ url, className }: Props) => {
	const Icon = socialIconByUrl(url)
	return <a className="text-foreground/50 transition-colors hover:text-foreground" href={url}><Icon className={cn(className)} /></a>
}

export default SocialIcon
