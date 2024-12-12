import { Outlet, redirect, useLoaderData } from "@remix-run/react";
import { ToastList } from "@/components/toast/toast-list";
import { ToastProvider } from "@/lib/context/toast.provider";
import { LoaderFunction, LoaderFunctionArgs } from "@remix-run/node";
import env from "@/config/env";
import { userSchema } from "@/lib/dto/user.dto";
import { useEffect } from "react";
import { useAuth } from "@/lib/context/auth.provider";
import { useUser } from "@/lib/context/user.provider";

export const loader = async ({ request }: LoaderFunctionArgs) => {
  const res = await fetch(`${env.BACKEND_URL}/auth/validate-session`, {
    credentials: "include",
    headers: request.headers,
  })

  if (!res.ok) {
    throw redirect("/auth/login")
  }

  const user = userSchema.parse(await res.json())

  if (!user.roles.some(role => role === "admin")) {
    throw redirect("/auth/login")
  }

  return { user }
}

export default function AdminLayout() {
  const { user } = useLoaderData<typeof loader>()

  const { setUser } = useUser()

  useEffect(() => {
    setUser(user)
  }, [user])

  return (
    <ToastProvider>
      <ToastList />
      <Outlet />
    </ToastProvider>
  );
}
