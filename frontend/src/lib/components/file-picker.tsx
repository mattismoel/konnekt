import type { InputHTMLAttributes } from "react";
import { cn } from "../clsx";

type Props = Omit<InputHTMLAttributes<HTMLInputElement>, "onChange"> & {
  onChange: (files: FileList | null) => void;
};

const FilePicker = ({ onChange, className, ...rest }: Props) => (
  <input
    {...rest}
    type="file"
    className={cn(
      "text-text file:rounded-full file:bg-zinc-100 file:px-6 file:py-2 file:font-medium file:text-zinc-900 file:transition-[filter] hover:file:brightness-80",
      className,
    )}
    onChange={(e) => onChange(e.currentTarget.files)}
  />
);

export default FilePicker;
