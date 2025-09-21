import { addHours, differenceInHours, differenceInMinutes, format } from "date-fns"
import type { Concert, Event } from "../event"
import { DATE_FORMAT } from "@/lib/date"
import { createContext, useContext } from "react"
import { Link } from "@tanstack/react-router"
import { cn } from "@/lib/clsx"
import LinkButton from "@/lib/components/ui/button/link-button"
import { generateGoogleCalendarEventUrl } from "@/lib/google-calendar"
import { FaCalendarPlus } from "react-icons/fa6"

type Props = {
	event: Event
}

type EventCalendarContext = {
	timelineStart: Date
	timelineEnd: Date
	timelineTotalMins: number;
}

const EventCalendarContext = createContext<EventCalendarContext | null>(null)

const useEventCalendarContext = () => {
	const ctx = useContext(EventCalendarContext)
	if (!ctx) throw new Error("No EventCalendarContext.Provider found")
	return ctx
}

const formatConcertDuration = (from: Date, to: Date) => {
	return `${format(from, "HH:mm")} - ${format(to, "HH:mm")}`
}

const padHour = (hour: number) => {
	return String(hour).padStart(2, "0")
}

const EventCalendar = ({ event }: Props) => {
	const firstConcert = event.concerts[0]
	const lastConcert = event.concerts[event.concerts.length - 1]

	const timelineStart = firstConcert.from.getMinutes() > 0 ? addHours(firstConcert.from, - 1) : firstConcert.from
	const timelineEnd = lastConcert.to.getMinutes() > 0 ? addHours(lastConcert.to, 1) : lastConcert.to
	const timelineTotalMins = differenceInMinutes(timelineEnd, timelineStart)

	const googleCalendarUrl = generateGoogleCalendarEventUrl(event)

	return (
		<EventCalendarContext.Provider value={{ timelineStart, timelineEnd, timelineTotalMins }}>
			<div className="flex flex-col gap-16">
				<div className="flex justify-between">
					<div>
						<h2 className="text-xl font-semibold text-heading">Program for {event.title}</h2>
						<span>{format(firstConcert.from, DATE_FORMAT)}</span>
					</div>
					<LinkButton
						variant="secondary"
						to={googleCalendarUrl.toString()}
						className="hidden sm:flex w-fit h-fit"
					>
						<FaCalendarPlus />Føj til kalender
					</LinkButton>
				</div>
				<div className="grid grid-cols-[64px_auto] gap-4">
					<Timeline />
					<ul className="relative">
						{event.concerts.map((concert) => <ConcertEntry key={concert.id} concert={concert} />)}
					</ul>
				</div>
				<LinkButton
					variant="secondary"
					to={googleCalendarUrl.toString()}
					className="sm:hidden w-full h-fit"
				>
					<FaCalendarPlus />Føj til kalender
				</LinkButton>

			</div>
		</EventCalendarContext.Provider>
	)
}


const Timeline = () => {
	const { timelineStart, timelineEnd } = useEventCalendarContext()
	const displayHours = differenceInHours(timelineEnd, timelineStart)

	return (
		<div className="flex flex-col">
			{[...Array(displayHours)].map((_, i) => (
				<div key={i} className="border-t border-zinc-800 h-32">
					{padHour(timelineStart.getHours() + i)}:00
				</div>
			))}
		</div>
	)
}

type ConcertEntryProps = {
	concert: Concert
}

const ConcertEntry = ({ concert }: ConcertEntryProps) => {
	const { timelineStart, timelineTotalMins } = useEventCalendarContext()
	const offsetRatio = differenceInMinutes(concert.from, timelineStart) / timelineTotalMins
	const durationMinutes = differenceInMinutes(concert.to, concert.from)

	return (
		<li
			style={{ top: `calc(${offsetRatio} * 100%)`, height: `calc(${durationMinutes / timelineTotalMins} * 100%)` }}
			className="group absolute w-full">
			<Link
				to="/artists/$artistId"
				params={{ artistId: concert.artist.id.toString() }}
				className={cn(
					"px-3 py-2 h-full flex justify-between bg-blue-950 border border-blue-900 rounded-sm transition-colors",
					"hover:bg-blue-900 hover:border-blue-700"
				)}
			>
				<span className="font-semibold text-blue-300 group-hover:underline">{concert.artist.name}</span>
				<span className="text-blue-400">{formatConcertDuration(concert.from, concert.to)}</span>
			</Link>
		</li>
	)
}

export default EventCalendar
