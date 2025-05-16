import LoginForm from '@/lib/components/login-form';
import { login, type loginForm } from '@/lib/features/auth';
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
			<LoginForm onSubmit={onSubmit} />
		</main>
	)
}
