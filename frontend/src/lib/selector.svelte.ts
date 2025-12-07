export class Selector<T extends { value: string }> {
	#max: number | undefined;
	entries = $state<T[]>([])
	selected = $state<T[]>([])

	constructor(entries: T[], selected?: T[], max?: number) {
		this.entries = entries

		if (selected) {
			this.selected = selected
		}

		this.#max = max
	}

	select = (value: string) => {
		const entry = this.entries.find(entry => entry.value === value)
		if (!entry) return

		// If max amount of entries are selected, return.
		if (this.#max && this.selected.length >= this.#max) return

		this.selected = [...this.selected, entry]
	}

	deselect = (value: string) => {
		this.selected = this.selected.filter(entry => entry.value !== value)
	}
}
