import { useEffect, useRef, useState } from "react";
import { BiX } from "react-icons/bi"
import { z } from "zod";
import { Button } from "@/components/ui/button";
import { useToast } from "@/lib/context/toast.provider";
import { cn } from "@/lib/utils";
import { Spinner } from "@/components/ui/spinner";

type Props = {
  show: boolean;
  name: string;

  onClose: () => void;
  onUploaded: (uploadUrl: string) => void;
}

export const ImageSelectorModal = ({ show, name, onClose, onUploaded }: Props) => {
  const { addToast } = useToast()
  const [imageFile, setImageFile] = useState<File | null>(null)
  const [previewURL, setPreviewURL] = useState("")
  const [uploading, setUploading] = useState(false)

  const inputRef = useRef<HTMLInputElement>(null)
  const dialogRef = useRef<HTMLDialogElement>(null)

  const handleSubmit = async () => {
    if (!imageFile) return
    setUploading(true)

    const formData = new FormData()

    formData.append(name, imageFile)

    const res = await fetch(`${window.ENV.BACKEND_URL}/events/coverImage`, {
      method: "post",
      credentials: "include",
      body: formData,
    })

    if (!res.ok) {
      addToast("Kunne ikke uploade billede...", "error")
      return
    }

    const url = z.string().url().parse(await res.text())

    console.log(url)

    onUploaded(url)
    setUploading(false)
    dialogRef.current?.close()
  }

  useEffect(() => {
    if (!imageFile) return

    setPreviewURL(URL.createObjectURL(imageFile))
  }, [imageFile])

  return (
    <dialog
      ref={dialogRef}
      open={show}
      className={cn(
        "w-full max-w-lg fixed top-1/2 -translate-y-1/2 border rounded-md",
      )}
    >
      <input
        hidden
        type="file"
        name={name}
        onChange={e => setImageFile(e.currentTarget?.files?.[0] || null)}
        ref={inputRef}
      />
      {/* HEADER */}
      <div className="p-2 border-b flex justify-between items-center">
        <h3>Vælg coverbillede</h3>
        <button onClick={onClose}>
          <BiX />
        </button>
      </div>

      {/* CONTENT */}
      <div className="bg-zinc-900 min-h-72 flex items-center justify-center">
        {previewURL ? (
          <img src={previewURL} alt="Cover preview" className="h-full w-full object-cover" />
        ) : (
          <button
            type="button"
            className="h-full w-full"
            onClick={() => inputRef.current?.click()}
          >
            Upload billede...
          </button>
        )}
      </div>

      {/* FOOTER */}
      <div className="p-2 flex justify-end">
        <Button
          type="button"
          disabled={previewURL === ""}
          onClick={handleSubmit}
        >
          {uploading ? <Spinner /> : "Upload"}
        </Button>
      </div>

    </dialog>
  )
}
