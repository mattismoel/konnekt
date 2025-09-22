import { useAuthContext } from '@/lib/context/auth'
import { loginFormSchema } from '@/lib/features/auth/auth'
import LoginForm from '@/lib/features/auth/components/login-form'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type z from 'zod'

export const Route = createFileRoute('/_app/auth/login')({
	component: RouteComponent,
})

function RouteComponent() {
	const navigate = useNavigate()
	const { loginMember } = useAuthContext()

	const handleSubmit = async (form: z.infer<typeof loginFormSchema>) => {
		await loginMember(form)
		navigate({ to: "/admin/events" })
	}

	return (
		<main className="min-h-svh flex justify-center items-center">
			<LoginForm onSubmit={handleSubmit} />
		</main>
	)
}
