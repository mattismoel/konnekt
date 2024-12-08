import { createContext, useContext, useState } from "react";

const TIMEOUT_DURATION_SECS = 6.0

export type Severity = "info" | "success" | "warning" | "error"

export type Toast = {
  id: string;
  severity: Severity;
  message: string;
}

type ToastContextProps = {
  toasts: Toast[];
  addToast: (message: string, severity: Severity) => void;
  removeToast: (id: string) => void;
}

const ToastContext = createContext<ToastContextProps | undefined>(undefined)

export const ToastProvider = ({ children }: { children: React.ReactNode }) => {
  const [toasts, setToasts] = useState<Toast[]>([])

  const timeoutMap = new Map<string, NodeJS.Timeout>([])

  const addToast = (message: string, severity: Severity = "info") => {
    const id = crypto.randomUUID()
    setToasts(prev => [...prev, { id, message, severity }])

    const timeout = setTimeout(() => removeToast(id), TIMEOUT_DURATION_SECS * 1000)
    timeoutMap.set(id, timeout)
  }

  const removeToast = (id: string) => {
    setToasts(prev => prev.filter(toast => toast.id !== id))
    clearTimeout(id)
  }

  return (
    <ToastContext.Provider value={{ toasts, addToast, removeToast }}>
      {children}
    </ToastContext.Provider>
  )
}

export const useToast = (): ToastContextProps => {
  const context = useContext(ToastContext);
  if (!context) {
    throw new Error("useToast must be used within a ToastProvider");
  }
  return context;
};
