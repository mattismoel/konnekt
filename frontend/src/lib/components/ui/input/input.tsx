import { cn } from "@/lib/clsx";
import { forwardRef, type InputHTMLAttributes } from "react";

type Props = InputHTMLAttributes<HTMLInputElement>

const Input = forwardRef<HTMLInputElement, Props>(({ className, ...rest }, ref) => (
	<input ref={ref} {...rest} className={cn("w-full px-4 py-2 border bg-background border-zinc-900 rounded-sm focus:outline-none", className)} />
))

export default Input
