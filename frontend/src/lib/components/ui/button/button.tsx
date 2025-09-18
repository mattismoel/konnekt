import { forwardRef, type ButtonHTMLAttributes } from "react";
import { baseClasses, variantClasses, type RootProps } from "./base";
import { cn } from "@/lib/clsx";

type Props = RootProps & ButtonHTMLAttributes<HTMLButtonElement>

const Button = forwardRef<HTMLButtonElement, Props>(({
	children,
	className,
	type = "button",
	variant = "primary",
	...rest
}, ref) => (
	<button ref={ref} type={type} {...rest} className={cn(baseClasses, variantClasses.get(variant), className)}>
		{children}
	</button>
))

export default Button
