import ToastList from '@/lib/components/toast-list'
import { ToastProvider } from '@/lib/context/toast'
import { createFileRoute, Outlet } from '@tanstack/react-router'

export const Route = createFileRoute('/_app/auth')({
  component: RouteComponent,
})

function RouteComponent() {
  return (
    <ToastProvider>
      <Outlet />
      <ToastList />
    </ToastProvider>
  )
}
