import { createFileRoute, Outlet, redirect } from '@tanstack/react-router'

export const Route = createFileRoute('/admin')({
	component: RouteComponent,
	beforeLoad: async ({ context: { auth } }) => {
		if (! await auth.isAuthenticated()) {
			throw redirect({ to: "/auth/login" })
		}
	}
})

function RouteComponent() {
	return <Outlet />
}
