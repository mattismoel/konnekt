type Src = {
	src: string;
	alt: string;
}

type Props = {
	srcs: Map<string, Src>
}

const SponsorDisplay = ({ srcs }: Props) => {
	return (
		<div className="p-8 sm:p-16 justify-center gap-12 flex flex-wrap bg-radial from-zinc-900 to-background to-[75%]">
			{Array.from(srcs).map(([href, { src, alt }]) => (
				<a key={href} aria-label={alt} href={href} className="w-20 sm:w-24 group flex items-center justify-center">
					<img src={src} alt={alt} loading="lazy" className="aspect-square md:brightness-60 group-hover:brightness-100 transition-[filter,scale]" />
				</a>
			))}
		</div>
	)
}

export default SponsorDisplay
