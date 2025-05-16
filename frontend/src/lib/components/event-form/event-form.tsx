import { FormProvider, useFieldArray, useForm, type UseFieldArrayReturn, type UseFormReturn } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { FaArrowsRotate, FaPlus, FaUpload } from "react-icons/fa6"
import { addMinutes, roundToNearestHours } from "date-fns"
import ConcertList from "./concert-list"
import { z } from "zod"
import { createEventForm, editEventForm, type Event } from "@/lib/features/event"
import type { Venue } from "@/lib/features/venue"
import FormField from "../form-field"
import ImagePreview from "../image-preview"
import Input from "../ui/input"
import Selector from "../ui/selector"
import Button from "../ui/button/button"
import LinkButton from "../ui/button/link-button"
import type { Artist } from "@/lib/features/artist"
import { createContext, useContext } from "react"
import Tiptap from "../tiptap/tiptap"

// We must include the fieldArray as part of the context, to not create another
// instance in children.
//
// This also means, that we are NOT to use the react-hook-form useFormContext,
// but rather the extended useEventFormContext!
type EventFormContext =
	UseFieldArrayReturn<EventForm> &
	UseFormReturn<EventForm> & {
		event?: Event;
		venues: Venue[]
		artists: Artist[]

		onAddConcert: () => void;
		onDeleteConcert: (index: number) => void
		disabled: boolean;
	}

const EventFormContext = createContext<EventFormContext | undefined>(undefined)

export const useEventFormContext = () => {
	const eventFormContext = useContext(EventFormContext)

	if (!eventFormContext) throw new Error("No provider for EventFormContext")

	return eventFormContext
}


const eventFormSchema = z.union([createEventForm, editEventForm])
type EventForm = z.infer<typeof eventFormSchema>

type Props = {
	event?: Event;
	venues: Venue[]
	artists: Artist[]
	disabled?: boolean;

	onSubmit: (form: EventForm) => void;
}

const EventForm = ({ event, venues, artists, disabled = false, onSubmit }: Props) => {
	const methods = useForm<EventForm>({
		defaultValues: {
			...event,
			venueId: event?.venue.id,
			concerts: event?.concerts.map(c => ({
				from: c.from,
				to: c.to,
				artistID: c.artist.id,
			}))
		},
		resolver: zodResolver(eventFormSchema),
	})

	const { control, formState: { defaultValues, errors }, setValue, getValues, handleSubmit } = methods;

	const fieldArrayMethods = useFieldArray({ control, name: "concerts" })
	const { fields, remove, append } = fieldArrayMethods


	const onAddConcert = () => {
		const prevEnd = fields.length > 0
			? getValues(`concerts.${fields.length - 1}.to`)
			: undefined

		const from = prevEnd ? addMinutes(prevEnd, 15) : roundToNearestHours(new Date())
		const to = addMinutes(from, 30)
		append({ artistID: 1, from, to })
	}

	const onDeleteConcert = (idx: number) => {
		remove(idx)
	}

	return (
		<FormProvider {...methods}>
			<EventFormContext.Provider value={{ ...methods, ...fieldArrayMethods, artists, venues, event, disabled, onAddConcert, onDeleteConcert }}>
				<form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-16">
					<FormField error={errors.image}>
						<ImagePreview src={event?.imageUrl} accept="image/jpeg,image/png" onChange={file => setValue("image", file)} />
					</FormField>

					<GeneralSection />

					<FormField error={errors.description}>
						<Tiptap content={defaultValues?.description} onChange={(html) => setValue("description", html)} />
					</FormField>

					<ConcertsSection />

					<Button className="w-full md:w-fit" type="submit"><FaUpload />Offentliggør</Button>
				</form>
			</EventFormContext.Provider>
		</FormProvider>
	)
}

const GeneralSection = () => {
	const { register, formState: { errors } } = useEventFormContext()

	return (
		<section>

			<h1 className="text-2xl font-heading font-bold mb-4">Generelt</h1>
			<div className="flex flex-col gap-4">
				<FormField error={errors.title}>
					<Input {...register("title")} placeholder="Eventtitel" />
				</FormField>

				<div className="flex gap-4">
					<FormField error={errors.ticketUrl}>
						<Input {...register("ticketUrl")} placeholder="Billet-URL" className="w-full" />
					</FormField>

					<VenueSelector />
				</div>
			</div>
		</section>
	)
}

const ConcertsSection = () => {
	const { fields } = useEventFormContext()
	return (
		<section>
			<h1 className="text-2xl font-bold font-heading mb-4">Koncerter</h1>
			<ConcertList>
				{fields.map((field, index) => (
					<ConcertList.Entry key={field.id} index={index} />
				))}
			</ConcertList>
		</section>
	)
}

const VenueSelector = () => {
	const { venues, disabled } = useEventFormContext()

	const {
		setValue,
		getValues,
		formState: { errors },
	} = useEventFormContext()

	return (
		<FormField error={errors.venueId}>
			<Selector
				onChange={e => setValue("venueId", parseInt(e.target.value))}
				defaultValue={venues.find(v => v.id === getValues("venueId"))?.id}
				placeholder="Vælg venue..."
				className="w-full"
			>
				{venues.map(v => (
					<option key={v.id} value={v.id}>{v.name}</option>
				))}
			</Selector>

			{!disabled && (
				<div className="flex gap-2">
					<Button variant="ghost" className="aspect-square h-full">
						<FaArrowsRotate />
					</Button>
					<LinkButton to="/admin/venues/create" className="aspect-square h-full">
						<FaPlus />
					</LinkButton>
				</div>
			)}
		</FormField>
	)
}

export default EventForm
