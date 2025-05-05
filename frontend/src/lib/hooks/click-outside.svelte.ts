import type { Action } from "svelte/action"

type ClickOutsideAction = Action<HTMLElement, any, {
	onclickoutside: (e: CustomEvent) => void,

}>

export const clickOutside: ClickOutsideAction = (element) => {
	const handleClick = (event: MouseEvent) => {
		if (!element) return

		// Return early if the event is default prevented.
		if (event.defaultPrevented) return

		// Return if the clicked element is contained within the node.
		if (element.contains(event.target as HTMLElement)) return

		const clickOutsideEvent = new CustomEvent("clickoutside")
		element.dispatchEvent(clickOutsideEvent)
	}

	document.addEventListener("click", handleClick, true)

	return {
		destroy: () => document.removeEventListener("click", handleClick, true)
	}

}
