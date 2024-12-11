import { useToast } from "@/lib/context/toast.provider"
import { ToastEntry } from "./toast";

export const ToastList = () => {
  const { toasts } = useToast();

  return (
    <div
      className={`z-50 fixed bottom-0 right-0 w-full max-w-lg p-8 
        flex flex-col gap-2
        ${toasts.length < 0 && "pointer-events-none"}`}
    >
      {toasts.map(toast => (
        <ToastEntry key={toast.id} toast={toast} />
      ))}
    </div>
  )
}
