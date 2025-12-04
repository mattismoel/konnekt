type StringKeys<T> = {
	[K in keyof T]: T[K] extends string ? K : never;
}[keyof T];

export class EntrySearcher<T> {
	search = $state("");
	#entries = $state<T[]>([]);
	#property: string;

	results = $derived.by(() => {
		if (this.search === "") {
			return this.#entries;
		}

		return this.#entries.filter((entry) => {
			const value = entry[this.#property] as string;
			return value.toLowerCase().includes(this.search.toLowerCase());
		});
	});

	constructor(entries: T[], property: StringKeys<T>) {
		this.#entries = entries;
		this.#property = property as string;
	}
}
