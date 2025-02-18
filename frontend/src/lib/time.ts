import { format } from "date-fns"

export const formatDateStr = (d: Date): string => {
	return format(d, "EEEE / dd.MM.yy")
}
