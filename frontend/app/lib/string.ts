/**
  * Capitalises an input string
  */
export const capitalise = (s: string): string => {
  if (s.length <= 0) return ""

  if (s.length === 1) return s.toUpperCase()

  return s.charAt(0).toUpperCase() + s.slice(1);
}
