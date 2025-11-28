import { cn } from "@/lib/clsx";
import type { HTMLAttributes } from "react";
import type { FieldError, FieldErrorsImpl, Merge } from "react-hook-form";

type Error =
  | Merge<
      FieldError,
      (
        | Merge<
            FieldError,
            FieldErrorsImpl<{
              value: string;
            }>
          >
        | undefined
      )[]
    >
  | FieldError
  | undefined;

type Props = HTMLAttributes<HTMLDivElement> & {
  error?: Error;
};

const FormField = ({ error, children, className }: Props) => (
  <div className={cn("flex w-full flex-col gap-2", className)}>
    <div className="flex w-full gap-4">{children}</div>

    {error && (
      <span className={cn("hidden text-sm text-red-500", { block: error })}>
        {error?.message}
      </span>
    )}
  </div>
);

export default FormField;
