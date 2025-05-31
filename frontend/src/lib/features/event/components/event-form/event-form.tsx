import { Controller, FormProvider, useFieldArray, useForm, type UseFieldArrayReturn, type UseFormReturn } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { FaArrowsRotate, FaPlus, FaUpload } from "react-icons/fa6"
import { addMinutes, roundToNearestHours } from "date-fns"
import ConcertList from "./concert-list"
import { createEvent, eventForm, updateEvent, type Event, type EventFormValues } from "@/lib/features/event/event"
import type { Venue } from "@/lib/features/event/venue"
import type { Artist } from "@/lib/features/artist/artist"
import { createContext, useContext } from "react"
import FormField from "@/lib/components/form-field"
import ImagePreview from "@/lib/components/image-preview"
import Tiptap from "@/lib/components/tiptap/tiptap"
import Button from "@/lib/components/ui/button/button"
import Input from "@/lib/components/ui/input"
import Selector from "@/lib/components/ui/selector"
import LinkButton from "@/lib/components/ui/button/link-button"
import { createSubmitHandler } from "@/lib/api"
import { useQueryClient } from "@tanstack/react-query"
import { useAuth } from "@/lib/context/auth"

// We must include the fieldArray as part of the context, to not create another
// instance in children.
//
// This also means, that we are NOT to use the react-hook-form useFormContext,
// but rather the extended useEventFormContext!
type EventFormContext =
	UseFieldArrayReturn<EventFormValues> &
	UseFormReturn<EventFormValues> & {
		event?: Event;
		venues: Venue[]
		artists: Artist[]

		onAddConcert: () => void;
		onDeleteConcert: (index: number) => void
	}

const EventFormContext = createContext<EventFormContext | undefined>(undefined)

export const useEventFormContext = () => {
	const eventFormContext = useContext(EventFormContext)

	if (!eventFormContext) throw new Error("No provider for EventFormContext")

	return eventFormContext
}

type Props = {
	event?: Event;
	venues: Venue[]
	artists: Artist[]
}


const EventForm = ({ event, venues, artists }: Props) => {
	const { hasPermissions } = useAuth()
	const queryClient = useQueryClient()

	const isEditable = hasPermissions(["edit:event"])

	const methods = useForm({
		disabled: !isEditable,
		defaultValues:
			event ? {
				title: event.title,
				description: event.description,
				venueId: event?.venue.id,
				image: undefined,
				ticketUrl: event.ticketUrl,
				concerts: event?.concerts.map(c => ({
					from: c.from,
					to: c.to,
					artistID: c.artist.id,
				})),
				isPublic: event.isPublic
			} : {
				isPublic: true,
				venueId: venues.find(v => v.name === "Posten")?.id
			},
		resolver: zodResolver(eventForm),
	})

	const { control, formState: { errors }, setValue, getValues, handleSubmit } = methods;

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



	const onSubmit = createSubmitHandler({
		action: async (form: EventFormValues) => {
			event ? await updateEvent(form, event.id) : await createEvent(form)
			await queryClient.invalidateQueries({ queryKey: ["events"] })
		},
		successMessage: event ? "Event opdateret" : "Event skabt",
		errorMessage: event ? "Kunne ikke redigére event" : "Kunne ikke lave event",
		navigateTo: "/admin/events",
	})

	return (
		<FormProvider {...methods}>
			<EventFormContext.Provider value={{ ...methods, ...fieldArrayMethods, artists, venues, event, onAddConcert, onDeleteConcert }}>
				<form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-16 @container">
					<FormField error={errors.image}>
						<ImagePreview disabled={!isEditable} src={event?.imageUrl} accept="image/jpeg,image/png" onChange={file => setValue("image", file)} />
					</FormField>

					<GeneralSection />

					<FormField error={errors.description}>
						<Controller
							control={control}
							name="description"
							render={({ field: { value, onChange } }) => (
								<Tiptap disabled={!isEditable} content={value} onChange={onChange} />
							)}
						/>
					</FormField>

					<ConcertsSection />

					{isEditable && (
						<Button className="w-full md:w-fit" type="submit"><FaUpload />Offentliggør</Button>
					)}
				</form>
			</EventFormContext.Provider>
		</FormProvider>
	)
}

const GeneralSection = () => {
	const { register, formState: { errors, disabled } } = useEventFormContext()
	const isEditable = !disabled

	return (
		<section>

			<h1 className="text-2xl font-heading font-bold mb-4">Generelt</h1>
			<div className="flex flex-col gap-4">
				<FormField error={errors.title}>
					<Input {...register("title")} placeholder="Eventtitel" />
				</FormField>

				<div className="flex flex-col @xl:flex-row gap-4">
					<FormField error={errors.ticketUrl}>
						<Input {...register("ticketUrl")} placeholder="Billet-URL" className="w-full" />
					</FormField>

					<VenueSelector />
				</div>
				{isEditable && (
					<FormField className="w-min">
						<label className="flex gap-2 items-center">
							<input type="checkbox" {...register("isPublic")} />
							Offentlig
						</label>
					</FormField>
				)}
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
	const { venues, formState: { disabled } } = useEventFormContext()
	const isEditable = !disabled

	const {
		setValue,
		getValues,
		formState: { errors },
	} = useEventFormContext()

	return (
		<FormField error={errors.venueId}>
			<Selector
				disabled={!isEditable}
				onChange={e => setValue("venueId", parseInt(e.target.value))}
				defaultValue={venues.find(v => v.id === getValues("venueId"))?.id}
				placeholder="Vælg venue..."
				className="w-full"
			>
				{venues.map(v => (
					<option key={v.id} value={v.id}>{v.name}</option>
				))}
			</Selector>

			{isEditable && (
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
