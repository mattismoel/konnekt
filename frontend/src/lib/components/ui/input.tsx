import { cn } from "@/lib/clsx";
import { forwardRef, type InputHTMLAttributes } from "react";

const Input = forwardRef<
  HTMLInputElement,
  InputHTMLAttributes<HTMLInputElement>
>(({ className, ...rest }, ref) => {
  return (
    <input
      ref={ref}
      {...rest}
      className={cn(
        "w-full rounded-sm border border-zinc-900 bg-background px-4 py-2 disabled:text-text/50",
        className,
      )}
    />
  );
});

export default Input;
