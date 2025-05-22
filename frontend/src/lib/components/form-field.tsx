import { cn } from "@/lib/clsx";
import type { HTMLAttributes } from "react";
import type { FieldError, FieldErrorsImpl, Merge } from "react-hook-form";

type Error = Merge<FieldError, (Merge<FieldError, FieldErrorsImpl<{
	value: string;
}>> | undefined)[]> | FieldError | undefined


type Props = HTMLAttributes<HTMLDivElement> & {
	error?: Error;
}

const FormField = ({ error, children, className }: Props) => (
	<>
		<div className={cn("flex gap-4 w-full", className)}>
			{children}
		</div>

		{error && (
			<span className={cn("text-sm hidden text-red-500", { "block": error })}>
				{error?.message}
			</span>
		)}
	</>
)

export default FormField
