import { cn } from "@/lib/clsx"
import { createLink } from "@tanstack/react-router"
import { forwardRef, type AnchorHTMLAttributes } from "react"
import { baseClasses, variantClasses, type RootProps } from "./base"

type Props = RootProps & AnchorHTMLAttributes<HTMLAnchorElement>

const LinkButton = createLink(
	forwardRef<HTMLAnchorElement, Props>(({
		variant = "primary",
		className,
		...rest
	}, ref) => {
		return <a ref={ref} {...rest} className={cn(baseClasses, variantClasses.get(variant), className)} />
	}))

export default LinkButton
