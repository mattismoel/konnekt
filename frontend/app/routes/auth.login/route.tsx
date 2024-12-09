import { ActionFunctionArgs, redirect } from "@remix-run/node";
import { LoginForm } from "./form";
import env from "~/config/env";
import { registerSchema } from "~/lib/auth/login-schema";

const LoginPage = () => {
  return (
    <main className="py-16 h-sub-nav flex justify-center">
      <LoginForm />
    </main>
  )
}

export const action = async ({ request }: ActionFunctionArgs) => {
  try {
    const formData = await request.formData()

    const data = registerSchema.parse(Object.fromEntries(formData))

    const res = await fetch(`${env.BACKEND_URL}/auth/login`, {
      method: "post",
      body: JSON.stringify(data),
      credentials: "include",
      headers: { "Content-Type": "application/json" }
    })

    if (res.ok) {
      return res
    }

  } catch (e) {
    console.error(e)
    return null
  }

}

export default LoginPage;
