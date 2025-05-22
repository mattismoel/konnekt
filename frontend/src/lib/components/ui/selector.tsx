import { cn } from "@/lib/clsx";
import { Children, forwardRef, type SelectHTMLAttributes } from "react";

type Props = SelectHTMLAttributes<HTMLSelectElement> & {
	placeholder?: string;
};

const Selector = forwardRef<HTMLSelectElement, Props>(({ placeholder = "Vælg...", defaultValue, children, className, ...rest }, ref) => (
	Children.count(children) > 0 && (
		<select
			ref={ref}
			{...rest}
			defaultValue={defaultValue || "placeholder"}
			className={cn('disabled:text-text/50 rounded-sm border border-zinc-900 bg-zinc-950', className)}
		>
			<option value="placeholder" disabled>{placeholder}</option>
			{children}
		</select>
	)
))

export default Selector
