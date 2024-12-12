import { Outlet } from "@remix-run/react";
import { Sidebar } from "./sidebar";
import { BottomBar } from "./bottom-bar";

const DashboardLayout = () => {
  return (
    <main className="min-h-sub-nav flex flex-col py-20">
      <div className="flex-1 grid grid-rows-1 grid-cols-1 px-8 md:px-32 sm:grid-cols-[256px_1fr] gap-4">
        <Sidebar className="hidden sm:block" />
        <Outlet />
        <BottomBar className="fixed bottom-0 sm:hidden" />
      </div>
    </main>
  )
}

export default DashboardLayout;
