import Logo from '@/lib/assets/logo'
import Footer from '@/lib/components/footer'
import Navbar from '@/lib/components/navbar/navbar'
import NavMenu from '@/lib/components/navmenu'
import { createFileRoute, Link, Outlet, useLocation } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import { BiMenu } from 'react-icons/bi'

export const Route = createFileRoute('/_app')({
  component: RouteComponent,
})

function RouteComponent() {
  const [expanded, setExpanded] = useState(false)

  const { pathname } = useLocation()

  useEffect(() => {
    setExpanded(false)
  }, [pathname])

  return (
    <>
      <Navbar>
        <Navbar.Header>
          <button className='block md:hidden' aria-label="Navigation menu button" onClick={() => setExpanded(true)}>
            <BiMenu className='text-2xl' />
          </button>
          <Link to="/">
            <Logo className='h-4' aria-label='Go to frontpage' />
          </Link>
        </Navbar.Header>

        <Navbar.RouteList>
          <Navbar.RouteEntry pathname="/events" name="Events" />
          <Navbar.RouteEntry pathname="/artists" name="Kunstnere" />
          <Navbar.RouteEntry pathname="/about" name="Om os" />
        </Navbar.RouteList>
      </Navbar>

      <NavMenu expanded={expanded} onClose={() => setExpanded(false)}>
        <NavMenu.RouteList>
          <NavMenu.RouteEntry to="/">Forside</NavMenu.RouteEntry>
          <NavMenu.RouteEntry to="/events">Events</NavMenu.RouteEntry>
          <NavMenu.RouteEntry to="/artists">Kunstnere</NavMenu.RouteEntry>
          <NavMenu.RouteEntry to="/about">Om os</NavMenu.RouteEntry>
        </NavMenu.RouteList>
      </NavMenu>

      <Outlet />

      <Footer />
    </>
  )
}
