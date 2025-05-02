import type { Snippet } from "svelte";

export type WithChildren<T> = T & { children: Snippet | undefined }
