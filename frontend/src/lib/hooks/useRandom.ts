import { useState } from "react"
import { randomIndex } from "../random"

	const [history, setHistory] = useState<number[]>(() => selection.length > 0 ? [randomIndex(selection.length)] : [])
export const useRandomIndex = <T,>(selection: T[], historySize: number = 2) => {

	const randomize = () => {
		if (selection.length === 0) return

		if (selection.length <= historySize) {
			const newIdx = randomIndex(selection.length)
			setHistory([newIdx])
			return
		}

		setHistory(prev => {
			let newIdx = randomIndex(selection.length)
			while (prev.includes(newIdx)) {
				newIdx = randomIndex(selection.length)
			}

			return [...prev, newIdx].slice(-historySize)
		})
	}

	const index = history[history.length - 1]

	return { index, randomize }
}
