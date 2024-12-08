import { Severity, Toast, useToast } from "~/lib/toast/toast"
import { BiInfoCircle, BiCheckCircle, BiError, BiErrorCircle, BiX } from "react-icons/bi";

type Props = {
  toast: Toast;
}

export const ToastEntry = ({ toast }: Props) => {
  const { removeToast } = useToast();
  const { id, message, severity } = toast

  const colorMap = new Map<Severity, string>([
    ["info", "bg-blue-950 text-blue-200 border-blue-800"],
    ["success", "bg-green-950 text-green-200 border-green-800"],
    ["warning", "bg-yellow-950 text-yellow-200 border-yellow-800"],
    ["error", "bg-red-950 text-red-200 border-red-800"],
  ]);


  return (
    <div
      className={`p-3 flex gap-4 items-center justify-between border rounded-sm 
      ${colorMap.get(severity)}`}
    >
      {severity === "info" && <BiInfoCircle />}
      {severity === "success" && <BiCheckCircle />}
      {severity === "warning" && <BiError />}
      {severity === "error" && <BiErrorCircle />}

      <span className="flex-1">{message}</span>
      <button type="button" onClick={() => removeToast(id)}>
        <BiX />
      </button
      >
    </div>
  )
}
