import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { LoginLoad, loginSchema } from "@/lib/dto/login-schema"

type Props = {
  onSubmit: (data: LoginLoad) => void;
}

export const LoginForm = ({ onSubmit }: Props) => {
  const { register, handleSubmit } = useForm<LoginLoad>({ resolver: zodResolver(loginSchema) })

  return (
    <form
      method="post"
      className="flex flex-col gap-2 max-w-lg"
      onSubmit={handleSubmit(onSubmit)}
    >
      <h1 className="text-2xl font-bold mb-4">Log in.</h1>
      <div className="space-y-2 mb-2">
        <Input {...register("email")} type="email" placeholder="Email" />
        <Input {...register("password")} type="password" placeholder="Adgangskode" />
      </div>
      <Button className="w-full">Login</Button>
    </form >
  )
}
