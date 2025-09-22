import { FaCircleUser } from "react-icons/fa6"
import Button from "./ui/button/button"
import { useRef, type InputHTMLAttributes } from "react"

type Props = Omit<InputHTMLAttributes<HTMLInputElement>, "type"> & {
	src?: string | undefined
}


const AvatarSelector = ({ src, ...rest }: Props) => {
	const inputRef = useRef<HTMLInputElement>(null)

	const handleOpenFileDialog = () => {
		inputRef.current?.click()
	}

	return (
		<div className="relative mb-4">
			{src ? <AvatarImage src={src} /> : <FaCircleUser className="text-9xl" />}
			<Button type="button" onClick={handleOpenFileDialog} className="absolute -bottom-4 -right-4 text-xs py-1 px-3 w-min">Vælg...</Button>
			<input ref={inputRef} {...rest} hidden type="file" />
		</div>
	)
}

type AvatarImageProps = {
	src?: string | undefined
}
const AvatarImage = ({ src }: AvatarImageProps) => (
	<img src={src} className="h-32 aspect-square rounded-full object-cover" />
)

export default AvatarSelector
