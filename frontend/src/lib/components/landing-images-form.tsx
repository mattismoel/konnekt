import { deleteLandingImage, uploadLandingImage, type Image } from "../features/content/landing"
import FilePicker from "./file-picker"
import { useState } from "react"
import Button from "./ui/button/button"
import { APIError, type ID } from "../api"
import { useToast } from "../context/toast"
import { useQueryClient } from "@tanstack/react-query"
import ImageList from "./image-list"
import { FaUpload } from "react-icons/fa6"
import { useAuth } from "../context/auth"

type Props = {
	images: Image[]
}

const LandingImagesForm = ({ images }: Props) => {
	const { hasPermissions } = useAuth()
	const { addToast } = useToast()
	const queryClient = useQueryClient()

	const isEditable = hasPermissions(["edit:content", "delete:content"])

	const [file, setFile] = useState<File | null>(null)

	const changeFile = (files: FileList | null) => {
		if (!files) return

		const file = files.item(0)

		if (!file) return
		setFile(file)
	}

	const handleUpload = async () => {
		if (!file) return

		try {
			await uploadLandingImage(file)
			await queryClient.invalidateQueries({
				queryKey: ["landing", "images"]
			})
			addToast("Billede uploadet")
		} catch (e) {
			if (e instanceof APIError) {
				addToast("Kunne ikke uploade billede", e.cause, "error")
				return
			}

			addToast("Kunne ikke uploade billede", "Noget gik galt...", "error")
			return
		}
	}

	const handleDelete = async (id: ID) => {
		if (!confirm("Er du sikker?")) return

		try {
			await deleteLandingImage(id)
			await queryClient.invalidateQueries({ queryKey: ["landing", "images"] })
			addToast("Billede slettet")
		} catch (e) {
			if (e instanceof APIError) {
				addToast("Kunne ikke slette billede", e.cause, "error")
				return
			}

			addToast("Kunne ikke slette billede", "Noget gik galt...", "error")
			return
		}
	}


	return (
		<form>
			{isEditable && (
				<div className="flex justify-between gap-4 mb-16">
					<FilePicker onChange={changeFile} />
					<Button type="button" variant="secondary" onClick={handleUpload}><FaUpload />Upload</Button>
				</div>
			)}

			<ImageList images={images} onDelete={handleDelete} />
		</form >
	)
}

export default LandingImagesForm
