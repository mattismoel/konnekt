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
			'text-text file:rounded-md file:bg-zinc-100 file:px-4 file:py-2 file:font-medium file:text-zinc-900',
			className
		)}
		onChange={(e) => onChange(e.currentTarget.files)}
	/>
)

export default FilePicker
