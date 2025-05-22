/**
 * @description Returns a random integer in range [0, max]
 * @author https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Math/random
 */
export const randomInt = (min: number, max: number) => {
  const minCeiled = Math.ceil(min)
  const maxFloored = Math.floor(max)

  return Math.floor(Math.random() * (maxFloored - minCeiled) + minCeiled)
}
