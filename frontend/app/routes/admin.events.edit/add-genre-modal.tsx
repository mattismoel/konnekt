import { forwardRef, useEffect, useState } from "react";
import { BiX } from "react-icons/bi";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

type Props = {
  show: boolean;
  existingGenres: string[];
  onClose: () => void;
  onSubmit: (genre: string) => void;
  className?: string;
}

export const AddGenreModal = ({ show, existingGenres, onClose, onSubmit, className }: Props) => {
  const [genre, setGenre] = useState("")
  const [exists, setExists] = useState(false)

  useEffect(() => {
    const exists = existingGenres.some(str => str.toLowerCase() === genre.toLowerCase())

    setExists(exists)
  }, [genre])


  const handleSubmit = () => {
    if (exists) alert("Genre already exists")
    onSubmit(genre)
  }

  return (
    <dialog
      open={show}
      className={cn(
        "rounded-sm p-4 fixed top-1/2 -translate-y-1/2",
        className
      )}
    >
      {/* CLOSE BUTTON */}
      <button type="button" onClick={onClose} className="absolute top-2 right-2">
        <BiX className="text-xl" />
      </button>

      {/* HEADER */}
      <div className="mb-4">
        <h3 className="text-xl font-bold">Tilføj ny genre</h3>
      </div>

      {/* CONTENT */}
      <div className="mb-4">
        <Input
          placeholder="Ny genre..."
          className={cn({ "text-red-500": exists })}
          onChange={(e) => setGenre(e.currentTarget.value)}
        />
      </div>

      {/* FOOTER */}
      <div>
        <Button
          type="button"
          className="w-full"
          onClick={handleSubmit}
        >
          Tilføj
        </Button>
      </div>
    </dialog>
  )
};
