import type { PropsWithChildren } from "react"

type Variant = "primary" | "secondary"

export type RootProps = PropsWithChildren<{
	variant?: Variant;
	disabled?: boolean;
}>

export const baseClasses = "flex justify-center items-center gap-4 px-5 py-2 rounded-sm font-medium border transition-colors w-full text-nowrap text-center"

export const variantClasses = new Map<Variant, string>([
	["primary", "bg-foreground text-background border-zinc-300 hover:bg-zinc-300 hover:border-zinc-400"],
	["secondary", "bg-foreground/10 text-heading/75 border-foreground/20 hover:bg-foreground/20 hover:text-heading"]
])

