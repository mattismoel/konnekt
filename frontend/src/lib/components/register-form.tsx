import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { registerForm } from "../features/auth";
import type { z } from "zod";
import Card from "./ui/card";
import ProfilePictureSelector from "./profile-picture-selector";
import Input from "./ui/input";
import Button from "./ui/button/button";
import FormField from "./form-field";

type Props = {
	onSubmit: (form: z.infer<typeof registerForm>) => void
}

const RegisterForm = ({ onSubmit }: Props) => {
	const { register, handleSubmit, setValue, formState: { errors } } = useForm({ resolver: zodResolver(registerForm) })

	return (
		<form onSubmit={handleSubmit(onSubmit)}>
			<Card className="max-w-lg">
				<Card.Header>
					<Card.Title>Tilmeld</Card.Title>
					<Card.Description>Her kan du tilmelde dig som medlem af foreningen Konnekt.</Card.Description>
				</Card.Header>

				<Card.Content className="gap-16">
					<div className="flex w-full justify-center">
						<FormField error={errors.profilePictureFile}>
							<div className="flex w-full justify-center">
								<ProfilePictureSelector
									onChange={(newFile) => setValue("profilePictureFile", newFile)}
								/>
							</div>
						</FormField>
					</div>

					<div className="flex flex-col gap-4">
						<div className="flex gap-4">
							<FormField error={errors.firstName}>
								<Input placeholder="Fornavn" {...register("firstName")} />
							</FormField>
							<FormField error={errors.lastName}>
								<Input placeholder="Efternavn" {...register("lastName")} />
							</FormField>
						</div>
						<FormField error={errors.email}>
							<Input type="email" placeholder="Email" {...register("email")} />
						</FormField>
					</div>

					<div className="flex flex-col gap-4">
						<FormField error={errors.password}>
							<Input type="password" placeholder="Adgangskode" {...register("password")} />
						</FormField>
						<FormField error={errors.passwordConfirm}>
							<Input type="password" placeholder="Gentag adgangskode" {...register("passwordConfirm")} />
						</FormField>
					</div>
				</Card.Content>

				<Card.Footer>
					<Button type="submit" className="w-full">Registrér</Button>
				</Card.Footer>
			</Card>
		</form>
	)
}

export default RegisterForm
