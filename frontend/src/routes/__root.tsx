import { createRootRoute, Link, Outlet } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <>
      <div className="p-2 flex gap-2">
        <Link to="/" >
          Home
        </Link>{' '}
        <Link to="/about" >
          About
        </Link>
      </div>
      <hr />
      <Outlet />
    </>
  )
})
