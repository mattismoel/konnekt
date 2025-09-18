import type { HTMLAttributes } from "react";
import { cn } from "../clsx";

type Props = HTMLAttributes<HTMLDivElement> & {
	mousePos: {
		left: number;
		top: number;
	}
	size?: "base" | "lg" | "xl"
}

const GlowingCursor = ({ mousePos, size = "base", className, ...rest }: Props) => {
	return (
		<div
			{...rest}
			className={cn("absolute -translate-x-1/2 -translate-y-1/2 h-72 blur-[128px] rounded-full aspect-square mix-blend-overlay bg-white opacity-100", {
				"h-96": size === "lg",
				"h-128": size === "xl",
			}, className)}
			style={{ top: `${mousePos.top}px`, left: `${mousePos.left}px` }}
		/>
	)
}

export default GlowingCursor

