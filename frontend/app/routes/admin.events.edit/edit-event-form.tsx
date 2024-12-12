import { BiPlus } from "react-icons/bi"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { CreateEditEventDTO, createEditEventSchema, EventDTO } from "@/lib/dto/event.dto"
import { GenreSelector } from "./genre-selector"
import { useEffect, useState } from "react"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"
import { DateTimePicker } from "@/components/ui/date-picker"
import { TipTapEditor } from "@/components/ui/tip-tap/editor"
import { AddGenreModal } from "./add-genre-modal"
import { useForm } from "react-hook-form"
import { zodResolver } from "@hookform/resolvers/zod"
import { ImageSelectorModal } from "./image-selector-modal"
import { CountryPicker } from "@/components/ui/country-picker"
import { FieldErrorList } from "@/components/ui/field-error-list"

type Props = {
  event: EventDTO | null
  genres: string[]
  onSubmit: (data: CreateEditEventDTO) => void
  className?: string;
}

export const EditEventForm = ({ event, genres, onSubmit, className }: Props) => {
  const [showCoverImageModal, setShowCoverImageModal] = useState(false)
  const [showGenreModal, setShowGenreModal] = useState(false)

  const [coverImageUrl, setCoverImageUrl] = useState<string | undefined>(event?.coverImageUrl)

  const [selectedGenres, setSelectedGenres] = useState(event?.genres || [])

  const { register, handleSubmit, setValue, formState: { errors } } = useForm<CreateEditEventDTO>({
    resolver: zodResolver(createEditEventSchema), defaultValues: {
      ...event,
      venue: event?.venue.name,
      genres: event?.genres,
    }
  })

  useEffect(() => {
    setValue("genres", selectedGenres)
  }, [selectedGenres, setValue])

  const isEdit = event !== null

  return (
    <form onSubmit={handleSubmit(d => console.log(d))} className={cn("", className)}>
      {/* COVER IMAGE */}
      <div className="relative aspect-video mb-4">
        {coverImageUrl ? (
          <img
            className="w-full h-full object-cover rounded-sm overflow-hidden"
            src={coverImageUrl}
            alt="Cover for event"
          />
        ) : (
          <div
            className="absolute flex justify-center items-center top-0 left-0 
              h-full w-full bg-zinc-900/50"
          >
            Vælg coverbillede...
          </div>
        )}
        <div className="absolute top-0 left-0 h-full w-full rounded-sm border border-white/50 mix-blend-overlay"></div>
        <Button
          type="button"
          className="absolute bottom-2 left-2"
          onClick={() => setShowCoverImageModal(true)} >
          {isEdit ? "Ændr" : "Vælg"}
        </Button>
      </div>

      {/* HIDDEN INPUT CONTAINING THE AUTOMATICALLY SET COVER IMAGE URL */}
      <input {...register("coverImageUrl")} type="hidden" value={coverImageUrl} />

      <ImageSelectorModal
        show={showCoverImageModal}
        name="coverImage"
        onClose={() => setShowCoverImageModal(false)}
        onUploaded={(url) => setCoverImageUrl(url)}
      />

      <FieldErrorList errors={[errors.coverImageUrl?.message]} />

      {/* GENERAL */}
      <div className="space-y-2 mb-6">
        <h3 className="text-xl font-bold">Generelt</h3>
        <div>
          <Label htmlFor="title">Titel</Label>
          <Input
            {...register("title")}
            defaultValue={event?.title}
            aria-invalid={errors.title ? true : false}
            errors={[errors.title?.message]}
          />
        </div>
        <div className="flex gap-2">
          <div>
            <Label>Venue</Label>
            <Input
              {...register("venue")}
              defaultValue={event?.venue.name || "Posten"}
              className="flex-1"
            />
          </div>
          <div>
            <Label>By</Label>
            <Input
              {...register("city")}
              defaultValue={event?.venue.city}
              className="flex-1"
            />
          </div>
          <div>
            <Label>Land</Label>
            <CountryPicker defaultValue="DK" />
          </div>
        </div>
        <FieldErrorList errors={[
          errors.venue?.message,
          errors.city?.message,
          errors.country?.message
        ]} />
        <div className="space-y-2">
          <div className="flex-1">
            <Label>Fra</Label>
            <DateTimePicker
              {...register("fromDate")}
              initialDate={event?.fromDate}
              className="w-full"
              placeholder="Start dato..."
              errors={[errors.fromDate?.message]}
            />
          </div>
          <div className="flex-1">
            <Label>Til</Label>
            <DateTimePicker
              {...register("toDate")}
              initialDate={event?.toDate}
              className="w-full"
              placeholder="Slut dato..."
              errors={[errors.toDate?.message]}
            />
          </div>
        </div>
      </div>

      {/* GENRES */}
      <div className="mb-8">
        <div className="flex justify-between items-center mb-2">
          <h2 className="text-xl font-bold mb-3">Genrer</h2>
          <button
            type="button"
            className="flex items-center gap-2"
            onClick={() => setShowGenreModal(true)}
          >
            <BiPlus />
            Tilføj
          </button>
        </div>
        <input {...register("genres")} type="hidden" />
        <GenreSelector
          genres={genres}
          selected={selectedGenres}
          onChange={(updatedGenres) => setSelectedGenres(updatedGenres)}
        />
        <FieldErrorList errors={[errors.genres?.message]} />
      </div>

      <AddGenreModal
        existingGenres={genres}
        show={showGenreModal}
        onClose={() => setShowGenreModal(false)}
        onSubmit={() => { }}
      />

      {/* DESCRIPTION EDITOR */}
      <div className="mb-4">
        <h2 className="text-2xl font-bold mb-4">Beskrivelse</h2>
        <input {...register("description")} type="hidden" />
        <TipTapEditor onChange={(value) => setValue("description", value)} />
        <FieldErrorList errors={[errors.description?.message]} />
      </div>

      {/* SUBMIT BUTTON */}
      <Button
        type="submit"
        className={cn(
          "w-[calc(100vw-2rem)] fixed bottom-4 left-4",
          "sm:static sm:bottom-auto sm:left-auto sm:w-full"
        )}
      >
        {isEdit ? "Redigér" : "Lav"}
      </Button>
    </form>
  )
}
