type Src = string

type Props = {
	srcs: Map<string, Src>
}

const SponsorDisplay = ({ srcs }: Props) => {
	return (
		<div className="p-8 sm:p-16 justify-center gap-12 flex flex-wrap bg-radial from-zinc-900 to-background to-[50%]">
			{Array.from(srcs).map(([href, src]) => (
				<a href={href} className="w-20 sm:w-24 group flex items-center justify-center">
					<img src={src} className="aspect-square md:brightness-75 group-hover:brightness-100 transition-[filter,scale]" />
				</a>
			))}
		</div>
	)
}

export default SponsorDisplay
