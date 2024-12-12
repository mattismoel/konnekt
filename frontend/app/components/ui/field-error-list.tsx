import { cn } from "@/lib/utils"

type Props = {
  errors?: (string | undefined)[]
}

export const FieldErrorList = ({ errors }: Props) => {
  if (!errors) return

  if (errors.length <= 0) return

  return (
    <ul className={cn("list-inside text-sm mt-2 text-red-600", { "list-disc": errors.length > 1 })}>
      {errors.map(error => error !== undefined && <li>{error}.</li>)}
    </ul>
  )
}
