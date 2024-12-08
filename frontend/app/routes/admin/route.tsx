import { Outlet } from "@remix-run/react";
import { ToastList } from "~/components/toast/toast-list";
import { ToastProvider } from "~/lib/toast/toast";

export default function AdminLayout() {
  return (
    <ToastProvider>
      <ToastList />
      <Outlet />
    </ToastProvider>
  );
}
