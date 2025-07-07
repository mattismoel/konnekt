import { cn } from "@/lib/clsx";
import type { PropsWithChildren } from "react";

type ID = string

export type Entry = {
	id: ID;
	value: string;
	name: string;
}

type Props = PropsWithChildren<{
	selected: boolean;
	disabled?: boolean;
	onToggle: () => void;
}>;

const Entry = ({ children, selected, disabled = false, onToggle }: Props) => {
	return (
		<button
			disabled={disabled}
			type="button"
			onClick={onToggle}
			className={cn(
				'flex w-full text-text/50 items-center gap-4 rounded-sm border border-transparent bg-zinc-950 p-2 hover:not-disabled:border-zinc-800',
				{ 'border-zinc-800 text-text bg-zinc-900': selected }
			)}
		>
			<ToggleBox selected={selected} />
			{children}
		</button>
	)
}

const ToggleBox = ({ selected }: { selected: boolean }) => (
	<div className="h-5 w-5 rounded-full border border-zinc-700 bg-zinc-800 p-1">
		<div
			className={cn('h-full w-full rounded-full bg-zinc-700', { 'bg-blue-500': selected })}
		></div>
	</div>
)

export default Entry
