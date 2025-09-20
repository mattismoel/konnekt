import { useState } from "react"
import { randomIndex } from "../random"

/**
 * @description Hook for generating random index without index repetition.
 * @param selection The array to generate indicies from.
 * @param historySize The buffer size of which previous items cannot be randomly selected. 
 */
export const useRandomIndex = <T,>(selection: T[], historySize: number = 2) => {
	const [indexHistory, setIndexHistory] = useState<number[]>([])

	const randomize = () => {
		if (selection.length === 0) return

		if (selection.length === 1) {
			setIndexHistory([0])
			return
		}

		setIndexHistory(prevHistory => {
			// If there are less items than the history can hold, we cannot enforce
			// uniqueness. Therefore we just return a random index, though making
			// sure the item is not directly repeated.
			if (selection.length <= historySize) {
				let newIdx = randomIndex(selection.length)
				const prevIdx = prevHistory[prevHistory.length - 1]
				while (newIdx === prevIdx) {
					newIdx = randomIndex(selection.length)
				}

				return [...prevHistory, newIdx].slice(-historySize)
			}

			let newIdx = randomIndex(selection.length)
			while (prevHistory.includes(newIdx)) {
				newIdx = randomIndex(selection.length)
			}

			return [...prevHistory, newIdx].slice(-historySize)
		})
	}

	const overrideIndex = (newIdx: number) => {
		if (newIdx < 0 || newIdx > selection.length - 1) {
			throw new Error("Override index out of bounds")
		}

		setIndexHistory(prev => [...prev, newIdx].slice(-historySize))
	}

	const index = indexHistory[indexHistory.length - 1]

	return { randomIndex: index, randomize, overrideIndex }
}
