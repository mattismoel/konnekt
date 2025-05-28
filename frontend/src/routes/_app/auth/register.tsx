import { APIError } from '@/lib/api'
import { useToast } from '@/lib/context/toast'
import { register, type registerForm } from '@/lib/features/auth/auth'
import RegisterForm from '@/lib/features/auth/components/register-form'
import { createFileRoute } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/_app/auth/register')({
  component: RouteComponent,
})

function RouteComponent() {
  const { addToast } = useToast()
  const onSubmit = async (form: z.infer<typeof registerForm>) => {
    try {
      await register(form)
      addToast("Du er nu registreret", "Vent på godkendelse af administrator")
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke registrere medlem", e.cause, "error")
        console.error(e)
        return
      }

      addToast("Kunne ikke registrere medlem", "Noget gik galt", "error")
      throw e
    }
  }

  return (
    <main className="min-h-svh px-auto flex justify-center items-center">
      <RegisterForm onSubmit={onSubmit} />
    </main>
  )
}
