import { type PropsWithChildren } from "react";
import Modal from "./modal";
import Button from "./button/button";
import { cn } from "@/lib/clsx";

type ID = string

export type Entry = {
	id: ID;
	value: string;
	name: string;
}

type Props = {
	title?: string;
	description?: string;
	show: boolean;

	entries: Entry[]
	selected: Entry[]

	onChange: (entries: Entry[]) => void;
	onClose: () => void;
}

const Picker = ({
	title = "Vælg...",
	description,
	show,

	entries,
	selected,

	onChange,
	onClose,
}: PropsWithChildren<Props>) => {
	const onToggle = (entryId: ID) => {
		const entry = entries.find(entry => entry.id === entryId)
		if (!entry) return

		const isSelected = selected.some(entry => entry.id === entryId)

		if (isSelected) {
			const newSelected = selected.filter(entry => entry.id !== entryId)
			onChange(newSelected)
			return
		}

		const newSelected = [...selected, entry]
		onChange(newSelected)
	}

	return (
		<Modal show={show} onClose={onClose}>
			<Modal.Header>
				<Modal.Title>{title}</Modal.Title>
				{description && <Modal.Description>{description}</Modal.Description>}
			</Modal.Header>

			<Modal.Content>
				<div className="flex flex-col gap-1">
					{entries.map(entry => (
						<Entry key={entry.id} selected={selected.some(selectedEntry => selectedEntry.id === entry.id)} onToggle={() => onToggle(entry.id)}>
							{entry.name}
						</Entry>
					))}
				</div>
			</Modal.Content>

			<Modal.Footer>
				<Button onClick={onClose}>Vælg</Button>
			</Modal.Footer>
		</Modal>
	)
}

type EntryProps = {
	selected: boolean;
	onToggle: () => void;
};

const Entry = ({ children, selected, onToggle }: PropsWithChildren<EntryProps>) => {
	return (
		<button
			type="button"
			onClick={onToggle}
			className={cn(
				'flex w-full text-text/50 items-center gap-4 rounded-sm border border-transparent bg-zinc-950 p-2 hover:border-zinc-800',
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

Picker.Entry = Entry

export default Picker
