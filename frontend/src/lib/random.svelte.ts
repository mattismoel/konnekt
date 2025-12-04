export const randomIndex = <T>(entries: T[]): number => {
	return Math.floor(Math.random() * entries.length);
};

export const randomEntry = <T>(entries: T[]): T | undefined => {
	return entries.at(randomIndex(entries));
};

export class Randomiser<T> {
	#historySize: number;
	#entries = $state<T[]>([]);
	#history = $state<number[]>([]);

	entry = $derived(this.#entries.at(this.#history.at(-1) ?? 0));
	idx = $derived(this.#history.at(-1));

	constructor(entries: T[], historySize = 3, startIdx = -1) {
		this.#entries = entries;
		this.#historySize = Math.min(entries.length, historySize);

		if (startIdx < 0) {
			this.randomise();
		}
	}

	randomise = () => {
		if (this.#entries.length === 0) {
			this.#history = [];
			return;
		}

		if (this.#entries.length <= this.#historySize) {
			let nextIdx = this.#history.at(-1);

			while (nextIdx === this.#history.at(-1)) {
				nextIdx = randomIndex(this.#entries);
			}

			this.#history = [...this.#history, nextIdx ?? 0].slice(-this.#historySize);
			return;
		}

		let nextIdx = this.#history.at(-1) ?? 0;

		while (this.#history.includes(nextIdx)) {
			nextIdx = randomIndex(this.#entries);
		}

		this.#history = [...this.#history, nextIdx].slice(-this.#historySize);
		console.log(this.#history);
	};
}
