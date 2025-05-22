import Button from "@/lib/components/ui/button/button"
import Card from "@/lib/components/ui/card"
import Input from "@/lib/components/ui/input"
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { login, loginForm, type LoginFormValues } from "../auth"
import { useNavigate } from "@tanstack/react-router"

const LoginForm = () => {
	const navigate = useNavigate()

	const { register, handleSubmit } = useForm({
		resolver: zodResolver(loginForm)
	})

	const onSubmit = async (form: LoginFormValues) => {
		try {
			await login(form)
			navigate({ to: "/admin/events" })
		} catch (e) {
			throw e
		}
	}

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
