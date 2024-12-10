import { BiPlus } from "react-icons/bi"
import { Input } from "~/components/ui/input"
import { Label } from "~/components/ui/label"
import { EventDTO } from "~/lib/event/event.dto"
import { GenreSelector } from "./genre-selector"
import { useEffect, useRef, useState } from "react"
import { Button } from "~/components/ui/button"
import { cn } from "~/lib/utils"
import { DateTimePicker } from "~/components/ui/date-picker"
import { FilePicker } from "~/components/ui/file-picker"
import { TipTapEditor } from "~/components/ui/tip-tap/editor"
import { AddGenreModal } from "./add-genre-modal"

type Props = {
  event: EventDTO | null
  genres: string[]
  className?: string;
}

export const EditEventForm = ({ event, genres, className }: Props) => {
  const [selectedGenres, setSelectedGenres] = useState(event?.genres || [])
  const [coverImageUrl, setCoverImageUrl] = useState<string | undefined>(event?.coverImageUrl)
  const [showGenreModal, setShowGenreModal] = useState(false)

  const changeCoverImage = (files: File | FileList) => {
    let file: File;

    if (files instanceof FileList) {
      file = files[0]
    } else {
      file = files
    }

    const newCoverImageUrl = URL.createObjectURL(file)

    setCoverImageUrl(newCoverImageUrl)
  }

  const isEdit = event !== null

  return (
    <form action="" className={cn("", className)}>
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
        <FilePicker
          name="coverImage"
          title="Cover-billede"
          className="absolute bottom-2 left-2"
          onChange={changeCoverImage}
        />
      </div>

      {/* GENERAL */}
      <div className="space-y-2 mb-6">
        <h3 className="text-xl font-bold">Generelt</h3>
        <div>
          <Label htmlFor="title">Titel</Label>
          <Input name="title" defaultValue={event?.title} />
        </div>
        <div className="flex gap-2">
          <div>
            <Label>Venue</Label>
            <Input name="venue" defaultValue={"Posten"} className="flex-1" />
          </div>
          <div>
            <Label>By</Label>
            <Input name="city" defaultValue={event?.address.city} className="flex-1" />
          </div>
        </div>
        <div className="space-y-2">
          <div className="flex-1">
            <Label>Fra</Label>
            <DateTimePicker className="w-full" placeholder="Start dato..." />
          </div>
          <div className="flex-1">
            <Label>Til</Label>
            <DateTimePicker className="w-full" placeholder="Slut dato..." />
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
        <GenreSelector
          genres={genres}
          selected={selectedGenres}
          onChange={(updatedGenres) => setSelectedGenres(updatedGenres)}
        />
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
        <TipTapEditor />
      </div>

      {/* SUBMIT BUTTON */}
      <Button
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
