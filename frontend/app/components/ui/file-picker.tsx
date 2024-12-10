import { ChangeEvent } from "react";
import { cn } from "~/lib/utils";

type Props = {
  name: string;
  title: string;
  multiple?: boolean;
  accept?: string;
  className?: string;
  onChange: (file: File | FileList) => void;
}

export const FilePicker = ({ name, title, multiple, accept, onChange, className }: Props) => {
  const changeFile = (e: ChangeEvent<HTMLInputElement>) => {
    const files = e.currentTarget.files

    if (!files) return

    if (files.length === 1) {
      onChange(files[0])
      return
    }

    onChange(files)
  }

  return (
    <div className={cn("flex flex-col gap-1", className)}>
      <input
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
}
