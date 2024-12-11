import { LoaderFunction } from "@remix-run/node";
import env from "@/config/env";
import { userSchema } from "../lib/dto/user.dto";
import { Outlet } from "@remix-run/react";

export const loader: LoaderFunction = async ({ request }) => {
  const res = await fetch(`${env.BACKEND_URL}/auth/validate-session`, {
    credentials: "include",
    headers: request.headers,
  })

  if (!res.ok) {
    throw new Error("Failed to validate session")
  }

  const user = userSchema.parse(await res.json())

  return { user }
}


type Props = {
  requiredRoles: string[]
}

export const ProtectedRoute = ({ requiredRoles }: Props) => {
  return (
    <Outlet />
  )

}
