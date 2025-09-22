import { useState, type ChangeEvent } from "react"
import { useForm } from "react-hook-form"
import { registerFormSchema } from "../auth"
import { zodResolver } from "@hookform/resolvers/zod"
import type z from "zod"
import AvatarSelector from "@/lib/components/avatar-selector"
import Input from "@/lib/components/ui/input/input"
import Button from "@/lib/components/ui/button/button"

type Props = {
	onSubmit: (form: z.infer<typeof registerFormSchema>) => void;
}

const RegisterForm = ({ onSubmit }: Props) => {
	const [avatarSrc, setAvatarSrc] = useState("")

	const { register, setValue, handleSubmit } = useForm({
		resolver: zodResolver(registerFormSchema)
	})

	const handleUpdateAvatar = async (e: ChangeEvent<HTMLInputElement>) => {
		const file = e.currentTarget.files?.item(0)
		if (!file) return

		setValue("avatarFile", file)
		const newSrc = URL.createObjectURL(file)
		setAvatarSrc(newSrc)
	}
	return (
		<form className="flex flex-col gap-12 max-w-sm" onSubmit={handleSubmit(onSubmit)}>
			<h1 className="font-semibold font-heading text-heading text-3xl">Tilmeld ny bruger</h1>

			<div className="flex justify-center">
				<AvatarSelector src={avatarSrc} onChange={handleUpdateAvatar} />
			</div>

			<div className="flex flex-col gap-4">
				<fieldset className="flex gap-2">
					<Input placeholder="Fornavn" {...register("firstName")} />
					<Input placeholder="Efternavn" {...register("lastName")} />
				</fieldset>

				<Input type="email" placeholder="Email" {...register("email")} />

				<fieldset className="flex flex-col gap-2">
					<Input type="password" placeholder="Adgangskode" {...register("password")} />
					<Input type="password" placeholder="Gentag adgangskode" {...register("passwordConfirm")} />
				</fieldset>
			</div>

			<Button type="submit">Registrér</Button>
		</form>
	)
}

export default RegisterForm
