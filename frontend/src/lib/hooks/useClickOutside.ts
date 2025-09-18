import { useEffect, type RefObject } from "react";

export const useClickOutside = <T extends HTMLDivElement>(ref: RefObject<T | null>, handler: (e: MouseEvent) => void) => {
	useEffect(() => {
		if (!ref) return

		const handleEvent = (e: MouseEvent) => {
			console.log("HANDLED EVENT")
			const target = e.target as Node
			if (!target || !target.isConnected) return


			const isOutside = ref.current && !ref.current.contains(target)

			if (isOutside) {
				handler(e)
			}
		}

		window.addEventListener("mouseup", handleEvent)
		return () => window.removeEventListener("mouseup", handleEvent)
	})
}
