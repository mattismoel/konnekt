import type { Snippet } from "svelte";

export type WithClass<T = {}> = T & {
	class?: string;
};

export type WithChildren<T = {}> = T & {
	children?: Snippet<[]>;
};
