export const randomIndex = <T>(entries: T[]): number => {
	if (entries.length === 0) throw new Error("Input entries array is empty");
	return Math.floor(Math.random() * entries.length);
};

export const randomEntry = <T>(entries: T[]): T => {
	const idx = randomIndex(entries);
	if (!idx) throw new Error("Input entries array is empty");
	return entries[idx];
};

export class Randomiser<T> {
	#historySize: number;
	#entries = $state<T[]>([]);
	history = $state<number[]>([]);

	idx = $derived(this.history.at(-1));
	entry = $derived.by(() => (this.idx !== undefined ? this.#entries.at(this.idx) : undefined));

	constructor(entries: T[], historySize = 3, startIdx = -1) {
		this.#entries = entries;
		this.#historySize = Math.min(entries.length, historySize);

		if (startIdx < 0) {
			this.randomise();
		}
	}

	randomise = () => {
		// If there is only one entry, and the first selection has been made, no need to randomise.
		if (this.#entries.length === 1 && this.history.length !== 0) return;

		// If there is fewer entries than the history size, just select a new random entry, disregarding history.
		if (this.#entries.length <= this.#historySize) {
			let nextIdx = randomIndex(this.#entries);
			while (nextIdx === this.history.at(-1)) {
				nextIdx = randomIndex(this.#entries);
			}

			this.history = [...this.history, nextIdx].slice(-this.#historySize);
			return;
		}

		let nextIdx = randomIndex(this.#entries);
		while (this.history.includes(nextIdx)) {
			nextIdx = randomIndex(this.#entries);
		}

		this.history = [...this.history, nextIdx].slice(-this.#historySize);
	};
}
