import { createContext, useContext, useState, type PropsWithChildren } from "react";

const DEFAULT_LIFETIME_MS = 5000

export type Severity = "info" | "warning" | "error"

export type Toast = {
	id: string;
	title: string;
	message?: string;
	severity: Severity;
}

type ToastContext = {
	toasts: Toast[]

	addToast: (title: string, message?: string, severity?: Severity, lifetimeMs?: number) => void;
	removeToast: (id: string) => void;
}

const ToastContext = createContext<ToastContext | undefined>(undefined)

export const ToastProvider = ({ children }: PropsWithChildren) => {
	const [toasts, setToasts] = useState<Toast[]>([])

	const addToast = (title: string, message?: string, severity: Severity = "info", lifetimeMs: number = DEFAULT_LIFETIME_MS) => {
		const id = crypto.randomUUID()

		const toast: Toast = { id, title, message, severity }
		setToasts(prev => [...prev, toast])

		setTimeout(() => {
			removeToast(id)
		}, lifetimeMs)
	};

	const removeToast = (id: string) => {
		setToasts(prev => prev.filter(toast => toast.id !== id))
	};

	return (
		<ToastContext.Provider value={{ toasts, addToast, removeToast }}>
			{children}
		</ToastContext.Provider>
	)
}


export const useToast = () => {
	const toastContext = useContext(ToastContext)
	if (!toastContext) throw new Error("ToastContext has no provider!")

	return toastContext
}
