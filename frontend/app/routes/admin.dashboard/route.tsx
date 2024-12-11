import { Outlet } from "@remix-run/react";
import { Sidebar } from "./sidebar";
import { useUser } from "@/lib/context/user.provider";

const DashboardLayout = () => {
  const { user } = useUser()
  return (
    <main className="min-h-sub-nav flex flex-col px-16 py-16">
      <h3 className="font-bold text-2xl mb-4">Godaften, {user?.firstName}.</h3>
      <div className="flex-1 grid grid-rows-1 grid-cols-[256px_1fr] gap-4">
        <Sidebar />
        <Outlet />
      </div>
    </main>
  )
}

export default DashboardLayout;
