import type { LiHTMLAttributes, PropsWithChildren } from "react"

type Props = {
	entries: string[]
}

const PillList = ({ entries, children }: PropsWithChildren<Props>) => (
	<ul className="flex items-center flex-wrap gap-2">
		{children}
		{entries.map(name => <Pill key={name}>{name}</Pill>)}
	</ul>
)

const Pill = ({ children, ...rest }: LiHTMLAttributes<HTMLLIElement>) => (
	<li
		{...rest}
		className="cursor-default h-10 w-fit rounded-full border border-zinc-800 bg-zinc-900 px-4 text-text/75 flex justify-center items-center"
	>
		{children}
	</li>
)

export default PillList
