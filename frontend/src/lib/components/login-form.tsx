import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { loginForm } from "../features/auth"
import Card from "./ui/card"
import Input from "./ui/input"
import Button from "./ui/button/button"
import type { z } from "zod"

type Props = {
	onSubmit: (form: z.infer<typeof loginForm>) => void
}

const LoginForm = ({ onSubmit }: Props) => {
	const { register, handleSubmit } = useForm({
		resolver: zodResolver(loginForm)
	})

	return (
		<form onSubmit={handleSubmit(onSubmit)}>
			<Card className="max-w-96">
				<Card.Header>
					<Card.Title>Login</Card.Title>
					<Card.Description>Her kan du logge ind som medlem på Konnekts dashboard.</Card.Description>
				</Card.Header>

				<Card.Content>
					<section className="flex flex-col gap-4">
						<Input
							type="email"
							placeholder="Email"
							{...register("email")}
						/>
						<Input
							type="password"
							placeholder="Adgangskode"
							{...register("password")}
						/>
					</section>
				</Card.Content>

				<Card.Footer>
					<Button className="w-full" type="submit">Login</Button>
				</Card.Footer>
			</Card>
		</form>
	)
}

export default LoginForm
