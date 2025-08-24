import { format } from "date-fns"
import type { Concert } from "./features/event/concert"
import type { Event } from "./features/event/event"

const GOOGLE_CALENDAR_DATETIME_FORMAT = "yyyyMMdd'T'HHmm'00'"

const googleCalendarUrl = "https://calendar.google.com/calendar/u/0/r/eventedit"

const formatGoogleCalendarDates = (first: Concert, last: Concert) => {
	const fromDateStr = format(first.from, GOOGLE_CALENDAR_DATETIME_FORMAT)
	const toDateStr = format(last.to, GOOGLE_CALENDAR_DATETIME_FORMAT)
	return `${fromDateStr}/${toDateStr}`
}

export const generateGoogleCalendarEventUrl = (event: Event): URL | null => {
	const firstConcert = event.concerts.at(0)
	const lastConcert = event.concerts.at(-1)

	if (!firstConcert || !lastConcert) return null

	const url = new URL(googleCalendarUrl)

	const dateParam = formatGoogleCalendarDates(firstConcert, lastConcert)
	const locationParam = `${event.venue.name}, ${event.venue.city}`
	url.searchParams.append("text", event.title)
	url.searchParams.append("dates", dateParam)
	url.searchParams.append("location", locationParam)
	return url
}
