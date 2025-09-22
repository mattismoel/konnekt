import Button from "@/lib/components/ui/button/button";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { loginFormSchema } from "../auth";
import type z from "zod";
import Input from "@/lib/components/ui/input/input";

type Props = {
	onSubmit: (form: z.infer<typeof loginFormSchema>) => void;
}

const LoginForm = ({ onSubmit }: Props) => {
	const { register, handleSubmit } = useForm({ resolver: zodResolver(loginFormSchema) })
	return (
		<form onSubmit={handleSubmit(onSubmit)} className="max-w-sm flex flex-col gap-8">
			<h1 className="font-heading text-heading font-semibold text-3xl">Log ind</h1>
			<div className="flex flex-col gap-2">
				<Input type="email" placeholder="Email" {...register("email")} />
				<Input type="password" placeholder="Adgangskode" {...register("password")} />
			</div>
			<Button type="submit">Log ind</Button>
		</form>
	)
}

export default LoginForm
