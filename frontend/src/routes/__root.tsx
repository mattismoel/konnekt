import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { TanStackRouterDevtoolsPanel } from '@tanstack/react-router-devtools'
import { TanstackDevtools } from '@tanstack/react-devtools'


import TanStackQueryDevtools from '../integrations/tanstack-query/devtools'

import type { QueryClient } from '@tanstack/react-query'
import Navbar from '@/lib/components/navbar/navbar'
import Footer from '@/lib/components/footer/footer'
import { useState } from 'react'
import NavMenu from '@/lib/components/navbar/nav-menu'

interface MyRouterContext {
	queryClient: QueryClient
}

const navEntries = [
	{ pathname: "/events", title: "Events" },
	{ pathname: "/artists", title: "Kunstnere" },
	{ pathname: "/about-us", title: "Om os" },
]


export const Route = createRootRouteWithContext<MyRouterContext>()({
	component: () => {
		const [navMenuOpen, setNavMenuOpen] = useState(false)

		return (
			<>
				<Navbar entries={navEntries} onOpenMenu={() => setNavMenuOpen(true)} menuOpen={navMenuOpen} />
				<NavMenu
					entries={[{ pathname: "/", title: "Forside" }, ...navEntries]}
					show={navMenuOpen}
					onClose={() => setNavMenuOpen(false)}
				/>
				<Outlet />
				<Footer />
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
