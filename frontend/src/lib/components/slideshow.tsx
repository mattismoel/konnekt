import { useEffect, useState, type ImgHTMLAttributes } from "react";
import { cn } from "../clsx";

// The amount of images to have passed, before the same image can be shown 
// again.
const HISTORY_SIZE = 3;

// The amount of seconds each image is to be shown.
const IMAGE_HOLD_SECS = 4.0

const rate = 1 / (1 / IMAGE_HOLD_SECS)

type Src = Pick<ImgHTMLAttributes<HTMLImageElement>, "alt" | "src">

type Props = {
	srcs: Src[]
}

const Slideshow = ({ srcs }: Props) => {
	let [indexHistory, setIndexHistory] = useState([0]);
	const currentIdx = indexHistory[indexHistory.length - 1]

	useEffect(() => {
		if (srcs.length <= 0) return

		const interval = setInterval(() => {
			setIndexHistory(prevHistory => {
				let nextIdx: number;

				do {
					nextIdx = Math.floor(Math.random() * srcs.length)
				} while (prevHistory.includes(nextIdx))

				const newHistory = [...prevHistory, nextIdx].slice(-HISTORY_SIZE);

				return newHistory
			})
		}, rate * 1000)

		return () => clearInterval(interval);
	}, [srcs.length])

	return (
		<div className="relative h-full w-full">
			{srcs.map(({ src, alt }, index) => (
				<img
					key={src}
					src={src}
					alt={alt}
					className={cn("absolute h-full w-full object-cover opacity-0 brightness-60 transition-opacity duration-1000", {
						"opacity-100": index === currentIdx
					})}
				/>
			))}
		</div>
	)
}

export default Slideshow
