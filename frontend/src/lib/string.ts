/**
 * @description Returns the input string with the first letter uppercased.
 *
 * @example "hello" => "Hello"
 * @example "some string" => "Some string"
 */
export const uppercaseFirstLetter = (s: string): string => {
  return s.charAt(0).toUpperCase() + s.slice(1);
};
