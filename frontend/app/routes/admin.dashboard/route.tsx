import { Outlet } from "@remix-run/react";
import { Sidebar } from "./sidebar";

const DashboardLayout = () => {
  return (
    <main className="min-h-sub-nav flex flex-col px-16 py-16">
      <div className="flex-1 grid grid-rows-1 grid-cols-[256px_1fr] gap-4">
        <Sidebar />
        <Outlet />
      </div>
    </main>
  )
}

export default DashboardLayout;
