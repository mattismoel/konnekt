import { createContext, useContext, useState } from "react"
import { z } from "zod"

import { Controller, FormProvider, useFieldArray, useForm, type UseFieldArrayReturn, type UseFormReturn } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"

import { FaPlus, FaTrash, FaUpload, FaPen } from "react-icons/fa"
import { artistForm, createArtist, socialUrlToIcon, updateArtist, type Artist, type ArtistFormValues } from "../artist"
import type { Genre } from "../genre"
import FormField from "@/lib/components/form-field"
import ImagePreview from "@/lib/components/image-preview"
import Button from "@/lib/components/ui/button/button"
import Input from "@/lib/components/ui/input"
import Tiptap from "@/lib/components/tiptap/tiptap"
import { trackIdFromUrl } from "@/lib/spotify"
import SpotifyPreview from "@/lib/components/spotify-preview"
import type { Entry } from "@/lib/components/ui/picker"
import PillList from "@/lib/components/pill-list"
import Picker from "@/lib/components/ui/picker"
import { createSubmitHandler } from "@/lib/api"
import { useQueryClient } from "@tanstack/react-query"
import { useAuth } from "@/lib/context/auth"

const internalSocialSchema = z.object({ value: z.string().url() })

const internalArtistFormSchema = artistForm
	.omit({ socials: true })
	.extend({ socials: internalSocialSchema.array() })

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
}

const ArtistForm = ({ artist, genres }: Props) => {
	const { hasPermissions } = useAuth()
	const isEditable = hasPermissions(["edit:artist"])

	const methods = useForm<InternalArtistForm>({
		disabled: !isEditable,
		defaultValues: {
			...artist,
			genreIds: artist?.genres.map(genre => genre.id) || [],
			socials: artist?.socials.map(s => ({ value: s }))
		},
		resolver: zodResolver(internalArtistFormSchema),
	})

	const { control, formState: { errors }, handleSubmit } = methods
	const fieldArrayMethods = useFieldArray({ control, name: "socials" })

	const queryClient = useQueryClient()

	const onSubmit = createSubmitHandler({
		navigateTo: "/admin/artists",
		successMessage: artist ? "Kunstner redigeret" : "Kunstner skabt",
		errorMessage: artist ? "Kunne ikke redigere kunstner" : "Kunne ikke skabe kunstner",
		action: async (form: InternalArtistForm) => {
			const socials: string[] = form.socials.map(({ value }) => value)
			const output: ArtistFormValues = { ...form, socials }

			const { data, success, error } = artistForm.safeParse(output)
			if (!success) {
				console.error(error)
				throw error
			}

			artist
				? await updateArtist(artist.id, data)
				: await createArtist(data)

			await queryClient.invalidateQueries({ queryKey: ["artists"] })
		},
	})

	return (
		<ArtistFormContext.Provider value={{ ...methods, ...fieldArrayMethods, artist, genres }}>
			<FormProvider {...methods}>
				<form className="flex flex-col gap-16" onSubmit={handleSubmit(onSubmit)}>
					<Controller
						control={control}
						name="image"
						render={({ field }) => (
							<FormField error={errors.image}>
								<ImagePreview
									{...field}
									src={artist?.imageUrl}
									accept="image/jpeg,image/png"
								// onChange={(file) => setValue("image", file)}
								/>
							</FormField>
						)}
					/>

					<GeneralSection />
					<SpotifySection />
					<GenreSection />
					<SocialsSection />

					{isEditable && (
						<Button type="submit" className="w-full md:w-fit"><FaUpload />Offentligør</Button>
					)}
				</form>
			</FormProvider>
		</ArtistFormContext.Provider>
	)
}

const GeneralSection = () => {
	const { control, register, formState: { errors } } = useArtistFormContext()

	return (
		<section >
			<h1 className="font-heading text-2xl font-bold mb-4">Generelt</h1>
			<div className="flex flex-col gap-4">
				<FormField error={errors.name}>
					<Input placeholder="Kunstnernavn" {...register("name")} />
				</FormField>

				<Controller
					control={control}
					name="description"
					render={({ field: { value, ...rest } }) => (
						<FormField error={errors.description}>
							<Tiptap {...rest} content={value} />
						</FormField>
					)}
				/>
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
	const { genres, control, formState: { errors, disabled } } = useArtistFormContext()
	const [showPicker, setShowPicker] = useState(false)

	const isEditable = !disabled

	const entries: Entry[] = genres.map(genre => ({
		id: genre.id.toString(),
		value: genre.id.toString(),
		name: genre.name,
	}))

	return (
		<section>
			<h1 className="font-bold font-heading mb-8 text-2xl">Genrer</h1>

			<Controller
				control={control}
				name="genreIds"
				render={({ field: { value, onChange, ...rest } }) => {
					const selectedEntries = entries.filter(e => value.includes(parseInt(e.value)))

					return (
						<>
							<FormField error={errors.genreIds}>
								<PillList entries={selectedEntries.map(entry => entry.name)}>
									{isEditable && (
										<Button
											variant="ghost"
											onClick={() => setShowPicker(true)}
											className="h-10 rounded-full px-4"
										>
											<FaPen />Vælg
										</Button>
									)}
								</PillList>
							</FormField>

							<Picker
								{...rest}
								title="Vælg genrer..."
								description="Her kan du vælge de genrer, som kunstneren associeres med."
								entries={entries}
								selected={selectedEntries}
								show={showPicker}
								onClose={() => setShowPicker(false)}
								onChange={(newEntries) => onChange(
									newEntries.map(e => parseInt(e.value))
								)}
							/>
						</>
					)
				}}
			/>
		</section>
	)
}

const SocialsSection = () => {
	const { fields, formState: { errors, disabled }, append } = useArtistFormContext()
	const isEditable = !disabled

	const [input, setInput] = useState("")

	const handleAdd = () => {
		append({ value: input })
		setInput("")
	}

	return (
		<section>
			<h1 className="font-heading font-bold text-2xl mb-4">Sociale medier</h1>

			{isEditable && (
				<div className="w-full flex gap-4 mb-8">
					<Input placeholder="URL..." value={input} onChange={e => setInput(e.target.value)} />
					<Button variant="secondary" onClick={handleAdd}><FaPlus /> Tilføj</Button>
				</div>
			)}

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
