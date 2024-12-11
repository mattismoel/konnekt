import { redirect, useNavigate } from "@remix-run/react";
import { createContext, useContext, useState } from "react";
import { LoginLoad, loginSchema } from "../dto/login-schema";
import { User, userSchema } from "../dto/user.dto";

type AuthContextType = {
  user: User | null;
  logIn: (data: LoginLoad) => void;
  logOut: () => void;
}

const AuthContext = createContext<AuthContextType | null>(null)

type Props = {
  children: React.ReactNode
}

export const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<User | null>(null);

  const navigate = useNavigate()

  const logIn = async (data: LoginLoad) => {
    try {
      const loginData = loginSchema.parse(data)

      const res = await fetch(`${window.ENV.BACKEND_URL}/auth/login`, {
        method: "post",
        body: JSON.stringify(loginData),
        credentials: "include",
        headers: { "Content-Type": "application/json" }
      })

      if (!res.ok) {
        throw new Error(`Could not login user: ${res.statusText}`)
      }

      const user = userSchema.parse(await res.json())
      setUser(user)
      navigate("/admin/dashboard/events")
    } catch (e) {
      throw new Error(`Could not login user: ${e}`)
    }
  }

  const logOut = async () => {
    const res = await fetch(`${window.ENV.BACKEND_URL}/auth/log-out`, {
      method: "post",
      credentials: "include"
    })

    if (!res.ok) {
      throw new Error("Could not log out")
    }

    navigate("/")
  }

  return (
    <AuthContext.Provider value={{ user, logIn, logOut }}>
      {children}
    </AuthContext.Provider>
  )
}

export const useAuth = () => {
  const authContext = useContext(AuthContext)

  if (!authContext) {
    throw new Error("useAuth has to be used within <AuthContext.Provider>")
  }

  return authContext
}
