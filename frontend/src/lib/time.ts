export const DATE_FORMAT = "EEEE / dd.MM.yy"
export const DATETIME_FORMAT = `${DATE_FORMAT}, HH:mm`

export const sleep = (durationMs: number): Promise<void> => {
  return new Promise(resolve => setTimeout(resolve, durationMs))
}
