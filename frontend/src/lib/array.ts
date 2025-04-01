export const removeDuplicates = <T extends { id: any }>(arr: T[]): T[] => {
	return arr.filter((value, index, self) => index === self.findIndex((t) => t.id === value.id))
}

/**
 * @description Returns a random value from the input array.
 * If a selected object is input, the returned element is ensured not to be the same as the selected.
 */
export const pickRandom = <T extends { id: any }>(arr: T[], selected?: T): T | undefined => {
	if (arr.length <= 0) return undefined

	if (!selected) {
		return arr[Math.floor(Math.random() * arr.length)]
	}

	let newEntry = selected

	while (newEntry.id === selected.id) {
		newEntry = arr[Math.floor(Math.random() * arr.length)]
	}

	return newEntry
}
