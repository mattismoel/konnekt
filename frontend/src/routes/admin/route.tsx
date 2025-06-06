import Logo from '@/lib/assets/logo'
import ContextMenu from '@/lib/components/context-menu'
import Navbar from '@/lib/components/navbar/navbar'
import NavMenu from '@/lib/components/navmenu'
import ToastList from '@/lib/components/toast-list'
import { AuthProvider, useAuth } from '@/lib/context/auth'
import { ToastProvider } from '@/lib/context/toast'
import { createFileRoute, Link, Navigate, Outlet, useLocation } from '@tanstack/react-router'
import { useEffect, useState } from 'react'
import { FaBars } from 'react-icons/fa'

export const Route = createFileRoute('/admin')({
  component: RouteComponent,
})

const AuthLayout = () => {
  const [expanded, setExpanded] = useState(false)
  const [showUserContext, setShowUserContext] = useState(false)

  const { member, refetch, handleLogout } = useAuth()
  const { pathname } = useLocation()

  useEffect(() => {
    setExpanded(false)
  }, [pathname])

  useEffect(() => {
    refetch()
  }, [])

  if (member === undefined) return "Loading..."

  if (member === null) return <Navigate to="/auth/login" />

  return <>
    <Navbar>
      <Navbar.Header>
        <button onClick={() => setExpanded(true)} className="md:hidden">
          <FaBars />
        </button>
        <Link to="/">
          <Logo className="h-4" />
        </Link>
      </Navbar.Header>

      <Navbar.Content>
        <Navbar.RouteList>
          <Navbar.RouteEntry pathname='/admin/events' name="Events" />
          <Navbar.RouteEntry pathname='/admin/artists' name="Kunstnere" />
          <Navbar.RouteEntry pathname='/admin/venues' name="Venues" />
          <Navbar.RouteEntry pathname='/admin/members' name="Medlemmer" />
          <Navbar.RouteEntry pathname='/admin/content' name="Indhold" />
        </Navbar.RouteList>
      </Navbar.Content>
      <button className="group" onClick={() => setShowUserContext(true)}>
        <img src={member.profilePictureUrl} alt="Profil" className="h-8 w-8 rounded-full object-cover outline outline-zinc-700 group-hover:outline-2" />
      </button>
      <ContextMenu show={showUserContext} onClose={() => setShowUserContext(false)}>
        <ContextMenu.LinkEntry
          to="/admin/members/$memberId"
          params={{ memberId: member.id.toString() }}
        >
          Redigér
        </ContextMenu.LinkEntry>
        <ContextMenu.Entry onClick={handleLogout}>
          Log ud
        </ContextMenu.Entry>
      </ContextMenu>
    </Navbar>

    <NavMenu expanded={expanded} onClose={() => setExpanded(false)}>
      <NavMenu.RouteList>
        <NavMenu.RouteEntry to="/admin/events">Events</NavMenu.RouteEntry>
        <NavMenu.RouteEntry to="/admin/artists">Kunstnere</NavMenu.RouteEntry>
        <NavMenu.RouteEntry to="/admin/venues">Venues</NavMenu.RouteEntry>
        <NavMenu.RouteEntry to="/admin/members">Medlemmer</NavMenu.RouteEntry>
        <NavMenu.RouteEntry to="/admin/content">Indhold</NavMenu.RouteEntry>
      </NavMenu.RouteList>
    </NavMenu>


    <div className="px-auto py-32">
      <Outlet />
      <ToastList />
    </div>
  </>
}

function RouteComponent() {
  return (
    <AuthProvider>
      <ToastProvider>
        <AuthLayout />
        <ToastProvider />
      </ToastProvider>
    </AuthProvider>
  )
}
