import AvatarImage from '@/lib/assets/avatar.png';
import { useRef, useState, type ChangeEvent, type InputHTMLAttributes } from 'react';

type Props = {
	src?: string;
	onChange: (newFile: File) => void;
	file?: File | null | undefined
	accept?: InputHTMLAttributes<HTMLInputElement>["accept"]
};

const ProfilePictureSelector = ({ src, accept = "image/jpeg,image/png", file, onChange }: Props) => {
	const [imgSrc, setImgSrc] = useState(() => file ? URL.createObjectURL(file) : src ? src : undefined)

	const ref = useRef<HTMLInputElement>(null)

	const onFileChange = (e: ChangeEvent<HTMLInputElement>) => {
		if (!e.target.files) return;
		const newFile = e.target.files.item(0);
		if (!newFile) return

		changeImage(newFile);
	};

	const changeImage = (newFile: File) => {
		setImgSrc(URL.createObjectURL(newFile))
		onChange(newFile)
	};

	return (
		<div className="relative w-fit">
			<input ref={ref} hidden accept={accept} type="file" onChange={onFileChange} />
			<img src={imgSrc || AvatarImage} alt="Profile" className="h-28 w-28 rounded-full object-cover" />
			<button
				type="button"
				onClick={() => ref.current?.click()}
				className="bg-text absolute right-0 bottom-0 translate-x-1/2 translate-y-1/2 rounded-sm px-2 py-1 text-sm text-zinc-950 shadow-sm"
			>
				Vælg
			</button>
		</div>
	)
}

export default ProfilePictureSelector
