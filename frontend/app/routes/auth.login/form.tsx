export const LoginForm = () => {
  return (
    <form method="post" className="flex flex-col gap-2 max-w-lg">
      <input type="email" name="email" required />
      <input type="password" name="password" required />
      <button>Login</button>
    </form>
  )
}
