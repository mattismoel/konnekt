import { useAuthContext } from '@/lib/context/auth'
import type { registerFormSchema } from '@/lib/features/auth/auth'
import RegisterForm from '@/lib/features/auth/components/register-form'
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type z from 'zod'

export const Route = createFileRoute('/_app/auth/register')({
	component: RouteComponent,
})

function RouteComponent() {
	const { registerMember } = useAuthContext()
	const navigate = useNavigate()

	const handleSubmit = async (form: z.infer<typeof registerFormSchema>) => {
		await registerMember(form)
		navigate({ to: "/admin/events" })
	}

	return (
		<main className="flex min-h-svh justify-center items-center">
			<RegisterForm onSubmit={handleSubmit} />
		</main>
	)
}
