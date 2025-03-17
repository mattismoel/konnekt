const DEFAULT_TOAST_LIFETIME_MS = 5000
export type Severity = "info" | "warning" | "error"


export type Toast = {
	id: string;
	title: string;
	message?: string;
	severity: Severity;
}

type Toaster = {
	addToast(title: string, message?: string, severity?: Severity, lifetimeMs?: number): void;
	removeToast(id: string): void;
	toasts: Toast[]
}

export const toaster = $state<Toaster>({
	toasts: [],
	addToast(title: string, message?: string, severity: Severity = "info", lifetimeMs: number = DEFAULT_TOAST_LIFETIME_MS) {
		const id = crypto.randomUUID()

		this.toasts = [...this.toasts, { id, title, message, severity }]

		setTimeout(() => {
			this.removeToast(id)
		}, lifetimeMs)
	},
	removeToast(id: string) {
		this.toasts = this.toasts.filter(toast => toast.id !== id)
	}
})
