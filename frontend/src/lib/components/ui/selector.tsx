import { cn } from "@/lib/clsx";
import { forwardRef, type SelectHTMLAttributes } from "react";

type Props = SelectHTMLAttributes<HTMLSelectElement> & {
  placeholder?: string;
};

const Selector = forwardRef<HTMLSelectElement, Props>(
  (
    { placeholder = "Vælg...", defaultValue, children, className, ...rest },
    ref,
  ) => (
    <select
      ref={ref}
      {...rest}
      defaultValue={defaultValue || "placeholder"}
      className={cn(
        "disabled:text-text/50 w-full rounded-full border border-zinc-800 bg-zinc-900 px-6 py-2",
        className,
      )}
    >
      <option value="placeholder" disabled>
        {placeholder}
      </option>
      {children}
    </select>
  ),
);

export default Selector;
