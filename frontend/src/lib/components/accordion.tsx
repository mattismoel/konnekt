import { useState, type PropsWithChildren } from "react"
import { cn } from "../clsx";
import { FaChevronDown } from "react-icons/fa6";

type Props = {
	title: string;
}

const Accordion = ({ title, children }: PropsWithChildren<Props>) => {
	const [expanded, setExpanded] = useState(false)

	return (
		<div className={cn("border border-zinc-900 rounded-sm group", { "expanded": expanded })}>
			<button
				type="button"
				className="flex items-center gap-4 px-8 py-4 bg-background w-full 
				text-left hover:bg-zinc-900 group-[.expanded]:bg-zinc-900 transition-colors"
				onClick={() => setExpanded(prev => !prev)}>
				<FaChevronDown className="group-[.expanded]:rotate-180 transition-transform" />
				<span className="font-medium">{title}</span>
			</button>

			<div className="px-8 py-8 hidden overflow-hidden group-[.expanded]:block cursor-default">
				<p className="prose prose-invert">{children}</p>
			</div>
		</div>
	)
}

export default Accordion
