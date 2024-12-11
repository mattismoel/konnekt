import { ChangeEvent, forwardRef } from "react";
import { cn } from "~/lib/utils";

type Props = {
  name: string;
  multiple?: boolean;
  accept?: string;
  className?: string;
  onChange: (file: FileList) => void;
}

export const FilePicker = forwardRef<HTMLInputElement, Props>(({ name, multiple, accept, onChange, className }: Props, ref) => {
  const changeFile = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.currentTarget.files

    if (!files) return

    onChange(files)
  }

  return (
    <div className={cn("flex flex-col gap-1", className)}>
      <input
        ref={ref}
        name={name}
        type="file"
        multiple={multiple}
        accept={accept}
        className="file:bg-foreground file:text-background file:px-2 file:py-1 
        file:border-none file:rounded-md file:font-medium"
        onChange={changeFile}
      />
    </div>
  )
})
