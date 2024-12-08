import { Button } from "~/components/ui/button"
import { Input } from "~/components/ui/input"

export const LoginForm = () => {
  return (
    <form method="post" className="flex flex-col gap-2 max-w-lg">
      <h1 className="text-2xl font-bold mb-4">Log in.</h1>
      <div className="space-y-2 mb-2">
        <Input type="email" name="email" placeholder="Email" required />
        <Input type="password" name="password" placeholder="Adgangskode" required />
      </div>
      <Button className="w-full">Login</Button>
    </form >
  )
}
