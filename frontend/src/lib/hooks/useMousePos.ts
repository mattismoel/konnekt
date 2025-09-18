import { useEffect, useState, type RefObject } from "react"
import { useScroll } from "./useScroll"

export const useMousePos = <T extends HTMLElement>(ref?: RefObject<T | null>) => {
	const [mousePos, setMousePos] = useState({ left: 0, top: 0 })

	const { diff } = useScroll()

	useEffect(() => {
		setMousePos(prev => ({
			left: prev.left + diff.x,
			top: prev.top + diff.y
		}))
	}, [diff.x, diff.y])

	const handleMouseMove = (e: Event) => {
		if (!(e instanceof MouseEvent)) return
		const element = ref?.current

		const newScrollX = element ? element.scrollLeft : window.scrollX
		const newScrollY = element ? element.scrollTop : window.scrollY

		if (!element) {
			setMousePos({
				left: e.clientX + newScrollX,
				top: e.clientY + newScrollY
			})
			return
		}

		const rect = element.getBoundingClientRect() ?? { left: 0, top: 0 };
		setMousePos({
			left: (e.clientX - rect.left) - newScrollX,
			top: (e.clientY - rect.top) - newScrollY,
		})
	}

	useEffect(() => {
		const target = ref?.current ?? window
		target.addEventListener("mousemove", handleMouseMove)

		return () => {
			target.removeEventListener("mousemove", handleMouseMove)
		}
	}, [ref, handleMouseMove])

	return mousePos
}
