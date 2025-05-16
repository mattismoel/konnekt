import { createContext, useContext, useState } from "react"
import { z } from "zod"

import { FormProvider, useFieldArray, useForm, type UseFieldArrayReturn, type UseFormReturn } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"

import { createArtistForm, editArtistForm, socialUrlToIcon, type Artist } from "../features/artist"
import type { Genre } from "../features/genre"
import { trackIdFromUrl } from "../spotify"

import FormField from "./form-field"
import Input from "./ui/input"
import Button from "./ui/button/button"
import Picker, { type Entry } from "./ui/picker"
import PillList from "./pill-list"
import ImagePreview from "./image-preview"
import Tiptap from "./tiptap/tiptap"
import SpotifyPreview from "./spotify-preview"

import { FaPlus, FaTrash, FaUpload, FaPen } from "react-icons/fa"

const artistForm = z.union([createArtistForm, editArtistForm])
type ArtistForm = z.infer<typeof artistForm>

const internalSocialSchema = z.object({ value: z.string().url() })

const internalArtistFormSchema = z.union([
	createArtistForm
		.omit({ socials: true })
		.extend({ socials: internalSocialSchema.array() }),
	editArtistForm
		.omit({ socials: true })
		.extend({ socials: internalSocialSchema.array() }),
])

type InternalArtistForm = z.infer<typeof internalArtistFormSchema>

type ArtistFormContext =
	UseFormReturn<InternalArtistForm> &
	UseFieldArrayReturn<InternalArtistForm> & {
		artist: Artist | undefined;
		genres: Genre[]
	}

const ArtistFormContext = createContext<ArtistFormContext | undefined>(undefined)

export const useArtistFormContext = () => {
	const ctx = useContext(ArtistFormContext)
	if (!ctx) throw new Error("No Provider for ArtistFormContext")
	return ctx
}

type Props = {
	artist?: Artist;
	genres: Genre[]

	onSubmit: (form: ArtistForm) => void;
}

const ArtistForm = ({ artist, genres, onSubmit }: Props) => {
	const methods = useForm<InternalArtistForm>({
		defaultValues: {
			...artist,
			genreIds: artist?.genres.map(genre => genre.id) || [],
			socials: artist?.socials.map(s => ({ value: s }))
		},
		resolver: zodResolver(internalArtistFormSchema),
	})

	const { control, formState: { errors }, register, watch, handleSubmit, setValue } = methods
	const fieldArrayMethods = useFieldArray({ control, name: "socials" })

	const submitForm = (form: InternalArtistForm) => {
		const output: ArtistForm = {
			...form,
			socials: form.socials.map(v => v.value)
		}

		const { data, success, error } = artistForm.safeParse(output)
		if (!success) {
			console.error(error)
			throw error
		}

		onSubmit(data)
	}

	return (
		<ArtistFormContext.Provider value={{ ...methods, ...fieldArrayMethods, artist, genres }}>
			<FormProvider {...methods}>
				<form className="flex flex-col gap-16" onSubmit={handleSubmit(submitForm)}>
					<FormField error={errors.image}>
						<ImagePreview
							src={artist?.imageUrl}
							accept="image/jpeg,image/png"
							onChange={(file) => setValue("image", file)}
						/>
					</FormField>

					<GeneralSection />
					<SpotifySection />
					<GenreSection />
					<SocialsSection />

					<Button type="submit" className="w-full md:w-fit"><FaUpload />Offentligør</Button>
				</form>
			</FormProvider>
		</ArtistFormContext.Provider>
	)
}

const GeneralSection = () => {
	const { register, setValue, formState: { errors, defaultValues } } = useArtistFormContext()

	return (
		<section >
			<h1 className="font-heading text-2xl font-bold mb-4">Generelt</h1>
			<div className="flex flex-col gap-4">
				<FormField error={errors.name}>
					<Input placeholder="Kunstnernavn" {...register("name")} />
				</FormField>

				<FormField error={errors.description}>
					<Tiptap
						content={defaultValues?.description}
						onChange={(html) => setValue("description", html)}
					/>
				</FormField>
			</div>
		</section>
	)

}

const SpotifySection = () => {
	const { register, watch } = useArtistFormContext()

	const trackId = trackIdFromUrl(watch("previewUrl"))

	return (
		<section>
			<h1 className="text-2xl font-bold font-heading mb-8">Spotify Preview</h1>
			<div className="flex flex-col gap-4">
				<Input placeholder="Spotify Preview-URL..." {...register("previewUrl")} />
				{trackId && <SpotifyPreview trackId={trackId} />}
			</div>
		</section>
	)
}

const GenreSection = () => {
	const { genres, setValue, formState: { errors }, watch } = useArtistFormContext()
	const [showPicker, setShowPicker] = useState(false)

	const entries: Entry[] = genres.map(genre => ({
		id: genre.id.toString(),
		value: genre.id.toString(),
		name: genre.name,
	}))

	const selected = entries.filter(entry =>
		watch("genreIds").includes(parseInt(entry.value))
	)

	return (
		<section>
			<h1 className="font-bold font-heading mb-8 text-2xl">Genrer</h1>

			<FormField error={errors.genreIds}>
				<PillList entries={selected.map(entry => entry.name)}>
					<Button variant="ghost" onClick={() => setShowPicker(true)} className="h-10 rounded-full px-4" ><FaPen />Vælg</Button>
				</PillList>
			</FormField>

			<Picker
				title="Vælg genrer..."
				description="Her kan du vælge de genrer, som kunstneren associeres med."
				entries={entries}
				selected={selected}
				show={showPicker}
				onClose={() => setShowPicker(false)}
				onChange={(selectedEntries) =>
					setValue("genreIds", selectedEntries.map(entry => parseInt(entry.value)))
				}
			/>
		</section>
	)
}

const SocialsSection = () => {
	const { fields, formState: { errors }, append } = useArtistFormContext()

	const [input, setInput] = useState("")

	return (
		<section>
			<h1 className="font-heading font-bold text-2xl mb-4">Sociale medier</h1>

			<div className="w-full flex gap-4 mb-8">
				<Input placeholder="URL..." value={input} onChange={e => setInput(e.target.value)} />
				<Button variant="secondary" onClick={() => append({ value: input })}><FaPlus /> Tilføj</Button>
			</div>

			<div className="flex flex-col gap-2">
				{fields.map((field, index) => (
					<FormField error={errors.socials}>
						<SocialMediaEntry key={field.id} index={index} />
					</FormField>
				))}
			</div>
		</section>
	)
}

const SocialMediaEntry = ({ index }: { index: number }) => {
	const { register, watch, remove } = useArtistFormContext()
	const Icon = socialUrlToIcon(watch(`socials.${index}.value`))

	return (
		<div className="w-full flex gap-4 items-center">
			<div className="relative w-full">
				<Icon className="absolute right-4 top-1/2 -translate-y-1/2 text-text/50" />
				<Input {...register(`socials.${index}.value`)} />
			</div>
			<button onClick={() => remove(index)} type="button" className=" text-text/50 hover:text-text h-full"><FaTrash /></button>
		</div>
	)
}

export default ArtistForm
