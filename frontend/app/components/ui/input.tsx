import * as React from "react"

import { cn } from "@/lib/utils"
import { FieldErrorList } from "./field-error-list"

type Props = {
  errors?: (string | undefined)[]
}

const Input = React.forwardRef<HTMLInputElement, React.ComponentProps<"input"> & Props>(
  ({ className, type, errors, ...props }, ref) => {
    return (
      <>
        <input
          type={type}
          className={cn(
            "flex h-9 w-full rounded-md border border-input bg-transparent px-3 py-1 text-base shadow-sm transition-colors file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
            className
          )}
          ref={ref}
          {...props}
        />
        <FieldErrorList errors={errors} />
      </>
    )
  }
)
Input.displayName = "Input"

export { Input }
