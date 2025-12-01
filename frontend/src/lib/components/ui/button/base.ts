import type { PropsWithChildren } from "react";

type Variant = "primary" | "secondary";

export type RootProps = PropsWithChildren & {
  variant?: Variant;
  disabled?: boolean;
};

export const baseClasses =
  "flex h-min w-fit items-center justify-center gap-3 rounded-full px-6 py-2 font-medium text-zinc-950 transition-[background,color,filter,border-color] disabled:opacity-50";

export const variantClasses = new Map<Variant, string>([
  ["primary", "bg-foreground text-text-dark hover:brightness-65"],
  [
    "secondary",
    "text-text-light-muted hover:text-text border border-foreground/10 bg-foreground/5 backdrop-blur-sm hover:border-foreground/30 hover:bg-foreground/10 text-text-light",
  ],
]);
