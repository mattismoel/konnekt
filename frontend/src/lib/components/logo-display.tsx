type Src = {
	name: string;
	src: string;
	href: string;
}

type Props = {
	srcs: Src[]
}

const LogoDisplay = ({ srcs }: Props) => (
	<div>
		<div className='w-full flex flex-wrap justify-center gap-12 p-8 bg-radial from-zinc-900 to-[70%] sm:p-16'>
			{srcs.map(({ name, src, href }) => (
				<a href={href} title={name}>
					<img src={src} alt={name} className='w-20 aspect-square hover:brightness-100 transition-[filter] sm:w-24 sm:brightness-75' />
				</a>
			))}
		</div>
	</div>
)

export default LogoDisplay;
