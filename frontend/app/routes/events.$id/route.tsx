import { LoaderFunctionArgs } from "@remix-run/node"
import { useLoaderData } from "@remix-run/react"
import env from "~/config/env"
import { eventSchema } from "~/lib/event/event.dto"

export const loader = async ({ params }: LoaderFunctionArgs) => {
  const id = parseInt(params.id || "")

  console.log(id)

  const res = await fetch(`${env.BACKEND_URL}/events/${id}`)

  if (!res.ok) {
    throw new Error(`Could not fetch for event with id ${id}: ${res.status}`)
  }

  const data = await res.json()

  const event = eventSchema.parse(data)
  return event
}

const EventPage = () => {
  const event = useLoaderData<typeof loader>()

  return (
    <main>Specific event</main>
  )
}

export default EventPage;
