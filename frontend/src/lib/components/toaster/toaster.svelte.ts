import { createContext } from "svelte";

type Severity = "info" | "warning" | "dangerous"

export type Toast = {
	id: string;
	text: string;
	severity: string;
}

export class ToasterState {
	toasts = $state<Toast[]>([])
	#idToTimeoutMap = new Map<string, number>();

	constructor() {
		$effect(() => {
			return () => this.#idToTimeoutMap.clear()
		})
	}

	add = (text: string, severity: Severity = "info", durationMs: number = 5000) => {
		const id = crypto.randomUUID()

		this.toasts = [...this.toasts, {
			id,
			text,
			severity,
		}]

		this.#idToTimeoutMap.set(id, setTimeout(() => this.remove(id), durationMs))
	}

	remove = (id: string) => {
		const timeout = this.#idToTimeoutMap.get(id)
		if (timeout) {
			clearTimeout(timeout)
		}

		this.toasts = this.toasts.filter(toast => toast.id !== id)
	}
}

export const [getToastContext, setToastContext] = createContext<ToasterState>()
