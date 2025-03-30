export const removeDuplicates = <T extends { id: any }>(arr: T[]): T[] => {
	return arr.filter((value, index, self) => index === self.findIndex((t) => t.id === value.id))
}
