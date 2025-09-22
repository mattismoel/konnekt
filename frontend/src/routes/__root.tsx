import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanstackDevtools } from '@tanstack/react-devtools'

import TanStackQueryDevtools from '../integrations/tanstack-query/devtools'

import type { QueryClient } from '@tanstack/react-query'
import type { AuthContext } from '@/lib/context/auth'

interface MyRouterContext {
	queryClient: QueryClient
	auth: AuthContext
}

export const Route = createRootRouteWithContext<MyRouterContext>()({
	component: () => {
		return (
			<>
				<Outlet />
				<TanstackDevtools
					config={{ position: 'bottom-left' }}
					plugins={[
						{ name: 'Tanstack Router', render: <TanStackRouterDevtoolsPanel /> },
						TanStackQueryDevtools,
					]}
				/>
			</>
		)
	}
})
