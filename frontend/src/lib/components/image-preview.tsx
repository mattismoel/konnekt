import { useState, type ImgHTMLAttributes, type InputHTMLAttributes } from "react";
import FilePicker from "./file-picker";

type Props = Omit<ImgHTMLAttributes<HTMLImageElement>, "onChange"> & {
	onChange: (file: File) => void;
	accept: InputHTMLAttributes<HTMLInputElement>['accept'];
	disabled?: boolean;
};

const ImagePreview = ({ src, accept, disabled, onChange, ...rest }: Props) => {
	let [url, setUrl] = useState(src);

	const updateImage = (file: File | null) => {
		if (!file) return;

		setUrl(URL.createObjectURL(file));
		onChange(file);
	};

	return (
		<div className="relative aspect-video w-full">
			<div className="absolute h-full w-full bg-gradient-to-t from-black/50"></div>
			<img {...rest} src={url} className="h-full w-full rounded-md border border-zinc-900 object-cover" />
			{!disabled && (
				<FilePicker
					accept={accept}
					onChange={(files) => updateImage(files?.[0] || null)}
					className="absolute bottom-4 left-4"
				/>
			)}
		</div>
	)
}

export default ImagePreview
