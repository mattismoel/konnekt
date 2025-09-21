import type { PropsWithChildren } from 'react'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'
import { format } from 'date-fns'
import type { IconType } from 'react-icons/lib'
import { FaCalendarDay, FaMapPin, FaMusic, FaTicket } from 'react-icons/fa6'
import { DATE_FORMAT } from '@/lib/date'
import { eventByIdQueryOpts } from '@/lib/features/events/query'
import { eventGenres } from '@/lib/features/artists/genre'
import type { Event } from '@/lib/features/events/event'
import LinkButton from '@/lib/components/ui/button/link-button'
import EventCalendar from '@/lib/features/events/components/event-calendar'

export const Route = createFileRoute('/_app/events/$eventId')({
	component: RouteComponent,
	loader: async ({ params: { eventId }, context: { queryClient } }) => {
		const id = parseInt(eventId)
		queryClient.ensureQueryData(eventByIdQueryOpts(id))
		return { eventId: id }
	}
})

function RouteComponent() {
	const { eventId } = Route.useLoaderData()
	const { data: event } = useSuspenseQuery(eventByIdQueryOpts(eventId))

	return (
		<main className="min-h-svh">
			<Header event={event} />
			<article className="px-responsive py-16 flex flex-col gap-16">
				<section dangerouslySetInnerHTML={{ __html: event.description }} />
				<section>
					<EventCalendar event={event} />
				</section>
			</article>
		</main >
	)
}

type HeaderProps = {
	event: Event
}

const Header = ({ event }: HeaderProps) => (
	<header className="relative h-[80svh]">
		<img src={event.imageUrl} className="h-full w-full object-cover" />
		<div className="fade-to-b-96 fade-background" />
		<div className="absolute bottom-8 left-0 px-responsive w-full">
			<h1 className="text-6xl font-heading text-heading font-semibold mb-6">{event.title}</h1>
			<div className="flex flex-col gap-8 sm:flex-row items-end">
				<ul className="w-full">
					<InfoEntry Icon={FaCalendarDay}>{format(event.concerts[0].from, DATE_FORMAT)}</InfoEntry>
					<InfoEntry Icon={FaMapPin}>{event.venue.name}, {event.venue.city} ({event.venue.country})</InfoEntry>
					<InfoEntry Icon={FaMusic}>{eventGenres(event).join(", ")}</InfoEntry>
				</ul>
				<LinkButton to={event.ticketUrl} className="w-full h-fit sm:w-fit"><FaTicket />Find billetter</LinkButton>
			</div>
		</div>
	</header>
)

type InfoEntryProps = PropsWithChildren<{
	Icon: IconType
}>

const InfoEntry = ({ Icon, children }: InfoEntryProps) => (
	<li className="flex gap-4 items-center">
		<Icon />
		{children}
	</li>
)
