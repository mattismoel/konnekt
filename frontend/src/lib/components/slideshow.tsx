import { useEffect, useState, type ImgHTMLAttributes } from "react";
import { cn } from "../clsx";

const RATE_MS = 1000 / (1 / 4);

type Src = Pick<ImgHTMLAttributes<HTMLImageElement>, "alt" | "src">

type Props = {
	srcs: Src[]
}

const Slideshow = ({ srcs }: Props) => {
	const [currentIdx, setCurrentIdx] = useState(0)


	useEffect(() => {
		if (srcs.length <= 0) return

		const interval = setInterval(() => {
			setCurrentIdx(prev => (prev + 1) % srcs.length)
			// Logic here...
		}, RATE_MS)

		return () => clearInterval(interval);
	}, [srcs.length])

	return (
		<div className="relative h-full w-full">
			{srcs.map(({ src, alt }, index) => (
				<img
					key={src}
					src={src}
					alt={alt}
					className={cn("absolute h-full w-full object-cover opacity-0 brightness-50 transition-opacity duration-1000", {
						"opacity-100": index === currentIdx
					})}
				/>
			))}
		</div>
	)
}

export default Slideshow
