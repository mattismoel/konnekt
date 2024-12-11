import { LoginForm } from "./form";
import { useAuth } from "@/lib/context/auth.provider";

const LoginPage = () => {
  const { logIn } = useAuth()

  return (
    <main className="py-16 h-sub-nav flex justify-center">
      <LoginForm onSubmit={logIn} />
    </main>
  )
}

export default LoginPage;
