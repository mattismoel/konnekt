import { register, type registerForm } from '@/lib/features/auth/auth'
import RegisterForm from '@/lib/features/auth/components/register-form'
import { createFileRoute } from '@tanstack/react-router'
import type { z } from 'zod'

export const Route = createFileRoute('/_app/auth/register')({
  component: RouteComponent,
})

function RouteComponent() {
  const onSubmit = async (form: z.infer<typeof registerForm>) => {
    try {
      await register(form)
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <main className="min-h-svh px-auto flex justify-center items-center">
      <RegisterForm onSubmit={onSubmit} />
    </main>
  )
}
