import LoginForm from '@/lib/features/auth/components/login-form';
import { createFileRoute } from '@tanstack/react-router'


export const Route = createFileRoute('/_app/auth/login')({
	component: RouteComponent,
})

function RouteComponent() {
	return (
		<main className="flex h-svh items-center justify-center px-auto">
			<LoginForm />
		</main>
	)
}
