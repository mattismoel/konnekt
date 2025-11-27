import { cn } from "../clsx";

type Props = {
	entries: {
		id: string,
		value: string
		ignorePrefix?: boolean;
	}[]
	selected: string;
	prefix?: string;

	className?: string;

	onSwitch: (newId: string) => void;
}

const Switch = ({ entries, selected, prefix, onSwitch, ...rest }: Props) => {
	return (
		<div className={cn("@container w-full flex justify-between items-center gap-8", rest.className)}>
			<div className={cn(
				"flex flex-col items-center justify-between gap-4 w-full border rounded-md border-transparent overflow-hidden  ",
				"@2xl:p-1 @2xl:flex-row @2xl:rounded-full @2xl:border-zinc-900",
				prefix && "w-full @2xl:pl-8",
			)}>

				<span className={cn(
					"text-center text-sm text-text/75 italic w-full whitespace-nowrap",
					"@2xl:text-left",
				)}>
					{prefix}</span>

				<div className="flex flex-col gap-1 w-full @2xl:flex-row">
					{entries.map(entry => (
						<button
							key={entry.id}
							type="button"
							onClick={() => onSwitch(entry.id)}
							title={entry.value}
							className={cn(
								"whitespace-nowrap border border-zinc-900 px-6 py-2 text-sm rounded-full transition-colors duration-100 hover:bg-zinc-900 hover:border-zinc-800",
								"before:invisible before:block before:h-0 before:overflow-hidden before:font-semibold before:content-[attr(title)]",
								selected === entry.id && "bg-zinc-100 font-semibold text-zinc-900 hover:bg-zinc-100 hover:border-zinc-100"
							)}
						>
							{entry.value}
						</button>
					))}
				</div>
			</div>
		</div>
	)
}

export default Switch
