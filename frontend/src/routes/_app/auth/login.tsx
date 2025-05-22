import { login, type loginForm } from '@/lib/features/auth/auth';
import LoginForm from '@/lib/features/auth/components/login-form';
import { createFileRoute, useNavigate } from '@tanstack/react-router'
import type { z } from 'zod';


export const Route = createFileRoute('/_app/auth/login')({
	component: RouteComponent,
})

function RouteComponent() {
	const navigate = useNavigate()
	const onSubmit = async (form: z.infer<typeof loginForm>) => {
		try {
			await login(form);
			navigate({ to: "/admin/events" })
		} catch (e) {
			throw e;
		}
	}

	return (
		<main className="flex h-svh items-center justify-center">
			<LoginForm />
		</main>
	)
}
