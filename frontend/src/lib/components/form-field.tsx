import { cn } from "@/lib/clsx";
import type { PropsWithChildren } from "react";
import type { FieldError } from "react-hook-form";

type Props = PropsWithChildren<{
	error?: FieldError | undefined
}>;

const FormField = ({ error, children }: Props) => (
	<div className="flex w-full flex-col gap-2">
		<div className="flex gap-4">
			{children}
		</div>

		{error && (
			<span className={cn("text-sm hidden text-red-500", { "block": error })}>
				{error?.message}
			</span>
		)}
	</div>
)

export default FormField
