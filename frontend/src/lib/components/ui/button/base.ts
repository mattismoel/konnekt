import type { PropsWithChildren } from "react";

type Variant = "primary" | "secondary" | "dangerous";

export type RootProps = PropsWithChildren & {
  variant?: Variant;
  disabled?: boolean;
};

export const baseClasses =
  "flex h-min w-fit items-center border justify-center gap-3 rounded-full px-6 py-2 font-medium text-zinc-950 transition-[background,color,filter,border-color] disabled:opacity-50";

export const variantClasses = new Map<Variant, string>([
  [
    "primary",
    "bg-foreground text-text-dark border-foreground hover:brightness-90",
  ],
  [
    "secondary",
    "text-text-light-muted hover:text-text border-foreground/10 bg-foreground/10 hover:border-foreground/30 hover:bg-foreground/40 hover:text-text-dark text-text-light",
  ],
  [
    "dangerous",
    "bg-red-950 border-red-900 text-red-500 hover:border-red-800 hover:bg-red-950 hover:text-red-500 hover:brightness-50",
  ],
]);
