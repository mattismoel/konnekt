import { format } from "date-fns"
import type { Event } from "./features/events/event"

const GOOGLE_CALENDAR_DATETIME_FORMAT = "yyyyMMdd'T'HHmm'00'"

const GOOGLE_CALENDAR_BASE_URL = "https://calendar.google.com/calendar/u/0/r/eventedit"

const formatGoogleCalendarDates = (fromDate: Date, toDate: Date) => {
	const fromDateStr = format(fromDate, GOOGLE_CALENDAR_DATETIME_FORMAT)
	const toDateStr = format(toDate, GOOGLE_CALENDAR_DATETIME_FORMAT)
	return `${fromDateStr}/${toDateStr}`
}

export const generateGoogleCalendarEventUrl = (event: Event): URL => {
	const firstConcert = event.concerts.at(0)
	const lastConcert = event.concerts.at(-1)

	if (!firstConcert) throw new Error("No concerts. Cannot create Google Calendar URL")

	const fromDate = firstConcert.from
	const toDate = lastConcert ? lastConcert.to : firstConcert.to

	const url = new URL(GOOGLE_CALENDAR_BASE_URL)

	const dateParam = formatGoogleCalendarDates(fromDate, toDate)
	const locationParam = `${event.venue.name}, ${event.venue.city}`

	url.searchParams.append("text", event.title)
	url.searchParams.append("dates", dateParam)
	url.searchParams.append("location", locationParam)

	return url
}
