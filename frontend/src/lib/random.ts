/**
 * @description Returns a random index of the input array length.
 */
export const randomIndex = (selectionCount: number) => Math.floor(Math.random() * selectionCount)

/**
 * @description Returns a random element from the input array.
 */
export const randomElement = <T,>(selection: T[]) => selection[randomIndex(selection.length)]
