import { useState } from 'react'
import { createFileRoute, Outlet } from '@tanstack/react-router'

import Footer from '@/lib/components/footer/footer'
import NavMenu from '@/lib/components/navbar/nav-menu'
import Navbar from '@/lib/components/navbar/navbar'

export const Route = createFileRoute('/_app')({
	component: RouteComponent,
})

const navEntries = [
	{ pathname: "/events", title: "Events" },
	{ pathname: "/artists", title: "Kunstnere" },
	{ pathname: "/about-us", title: "Om os" },
]

function RouteComponent() {
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
		</>
	)
}
