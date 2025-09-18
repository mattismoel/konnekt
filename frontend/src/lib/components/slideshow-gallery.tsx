import { Children, forwardRef, useEffect, useRef, useState, type ButtonHTMLAttributes, type PropsWithChildren } from "react";
import { useScroll } from "../hooks/useScroll";
import { cn } from "../clsx";
import { FaChevronLeft, FaChevronRight } from "react-icons/fa6";

const SlideshowGallery = ({ children }: PropsWithChildren) => {
	const ref = useRef<HTMLDivElement>(null)
	const { scroll } = useScroll(ref)
	const [scrollSize, setScrollSize] = useState(0)
	const [scrollIdx, setScrollIdx] = useState(0)

	const childCount = Children.count(children)

	useEffect(() => {
		if (!ref.current) return
		ref.current.scrollTo({ left: 0 })
		setScrollIdx(0)
	}, [])

	useEffect(() => {
		if (!ref.current || childCount === 0) return
		const newScrollSize = ref.current.scrollWidth / childCount
		setScrollSize(newScrollSize)
	}, [childCount])

	useEffect(() => {
		if (!ref.current || scrollSize === 0) return
		const newScrollIdx = Math.round(ref.current.scrollLeft / scrollSize)
		setScrollIdx(newScrollIdx)
	}, [scroll.x, scrollSize])

	const scrollToIndex = (idx: number) => {
		if (!ref.current || scrollSize === 0) return
		const clampedIdx = Math.max(0, Math.min(idx, childCount - 1))
		ref.current.scrollTo({
			left: clampedIdx * scrollSize,
			behavior: "smooth"
		})
	}

	const scrollInDirection = (direction: number) => {
		scrollToIndex(scrollIdx + direction)
	};

	return (
		<div className="flex flex-col gap-8">
			<div className="relative">
				<Button
					onClick={() => scrollInDirection(-1)}
					disabled={scrollIdx === 0}
					className="top-1/2 left-0 -translate-x-16 -translate-y-1/2">
					<FaChevronLeft />
				</Button>
				<Button
					onClick={() => scrollInDirection(1)}
					disabled={scrollIdx === childCount - 1}
					className="top-1/2  right-0 translate-x-16 -translate-y-1/2">
					<FaChevronRight />
				</Button>

				<div ref={ref} className="flex gap-4 overflow-x-scroll snap-x snap-mandatory scrollbar-none">
					{Children.map(children, (child, i) => (
						<div key={i} className="snap-center shrink-0 w-full">
							{child}
						</div>
					))}
				</div>
			</div>
			{childCount > 1 && <Dots amount={childCount} activeIdx={scrollIdx} />}
		</div>
	)
}

type DotsProps = {
	amount: number
	activeIdx: number
}

const Dots = ({ amount, activeIdx }: DotsProps) => (
	<div className="w-full flex items-center min-h-2 gap-1.5 justify-center">
		{[...Array(amount)].map((_, i) => (
			<div className={cn(
				"h-1.5 w-1.5 bg-zinc-800 rounded-full transition-[width,margin] ease-out",
				i === activeIdx && "bg-heading w-4 mx-1 first:ml-0 last:mr-0")} />
		))}
	</div>
)

const Button = forwardRef<HTMLButtonElement, ButtonHTMLAttributes<HTMLButtonElement>>(({ className, ...rest }, ref) => (
	<button
		ref={ref}
		{...rest}
		className={cn(
			"hidden absolute md:block px-3 py-2 bg-zinc-950 border border-zinc-900 rounded-sm transition-opacity cursor-pointer",
			"hover:border-zinc-800 hover:bg-zinc-900 disabled:opacity-0 disabled:cursor-default",
			className)}
	/>
))

export default SlideshowGallery
