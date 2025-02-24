import { createRootRoute, Link, Outlet } from "@tanstack/react-router";
import { Navbar } from "../components/navbar";
import { Footer } from "../components/footer";

export const Route = createRootRoute({
  component: () => (
    <>
      <Navbar />
      <main className="min-h-svh">
        <Outlet />
      </main>
      <Footer />
    </>
  )
})
