import { cn } from "@/lib/clsx";
import type { ButtonHTMLAttributes, PropsWithChildren } from "react"

const Toolbar = ({ children }: PropsWithChildren) => (
	<div className="flex gap-4 overflow-hidden rounded-t-md border border-zinc-800 bg-zinc-950">
		{children}
	</div>
)

const ActionGroup = ({ children }: PropsWithChildren) => (
	<div className="flex divide-x divide-zinc-700 overflow-hidden bg-zinc-900">
		{children}
	</div>
)

type ActionButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
	active?: boolean | undefined;
}

const ActionButton = ({ children, active, ...rest }: ActionButtonProps) => (
	<button
		{...rest}
		type="button"
		className={cn(
			'disabled:text-text/50 text-text/50 px-4 py-2 text-sm hover:bg-zinc-800 disabled:hover:bg-zinc-800',
			{ 'active bg-zinc-800 text-text': active }
		)}
	>
		{children}
	</button>
)

ActionGroup.Button = ActionButton

Toolbar.ActionGroup = ActionGroup

export default Toolbar
