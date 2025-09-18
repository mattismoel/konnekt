import { useEffect, useState, type RefObject } from "react"

export const useScroll = <T extends HTMLElement>(ref?: RefObject<T | null>) => {
	const [scroll, setScroll] = useState({ x: 0, y: 0 })
	const [diff, setDiff] = useState({ x: 0, y: 0 })
	const [direction, setDirection] = useState({ x: 0, y: 0 })

	useEffect(() => {
		const target = ref?.current ?? window

		const handleScroll = () => {
			setScroll((prev) => {
				const newScroll = (ref && ref.current)
					? ({ x: ref.current.scrollLeft, y: ref.current.scrollTop })
					: ({ x: window.scrollX, y: window.scrollY })


				const diffX = newScroll.x - prev.x
				const diffY = newScroll.y - prev.y

				setDiff({ x: diffX, y: diffY })

				setDirection({
					x: diffX === 0 ? 0 : diffX > 0 ? 1 : -1,
					y: diffY === 0 ? 0 : diffY > 0 ? 1 : -1
				})

				return newScroll
			})
		}

		target.addEventListener("scroll", handleScroll)
		return () => target.removeEventListener("scroll", handleScroll)
	}, [ref])

	return { scroll, direction, diff }
}
