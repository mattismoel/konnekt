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
        "disabled:text-text/50 w-full rounded-full border border-zinc-800 bg-zinc-900 px-6 py-2",
        className,
      )}
    />
  );
});

export default Input;
