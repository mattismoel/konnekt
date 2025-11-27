import { useState, type PropsWithChildren } from "react"
import { cn } from "../clsx";
import { FaChevronDown } from "react-icons/fa6";

type Props = {
	title: string;
}

const Accordion = ({ title, children }: PropsWithChildren<Props>) => {
	const [expanded, setExpanded] = useState(false)

	return (
		<div className={cn("border border-zinc-800 rounded-sm overflow-hidden group", { "expanded": expanded })}>
			<button
				type="button"
				className="text-base flex items-center gap-6 px-8 py-3 w-full bg-zinc-900 
				text-left hover:bg-zinc-800 group-[.expanded]:bg-zinc-800 transition-colors"
				onClick={() => setExpanded(prev => !prev)}>
				<FaChevronDown className="group-[.expanded]:rotate-180 transition-transform shrink-0 text-sm" />
				<span className="font-medium">{title}</span>
			</button>

			<div className="px-8 py-8 hidden overflow-hidden group-[.expanded]:block cursor-default">
				<p className="prose prose-invert leading-loose">{children}</p>
			</div>
		</div>
	)
}

export default Accordion
