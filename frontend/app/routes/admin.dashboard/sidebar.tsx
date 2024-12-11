import { Card, CardContent, CardHeader } from "@/components/ui/card"
import { SidebarEntry } from "./sidebar-entry"

export const Sidebar = () => {
  return (
    <Card>
      <CardHeader>
        <h3 className="font-black text-xl mb-8">KONNEKT &reg;</h3>
      </CardHeader>
      <CardContent>
        <ul className="flex flex-col gap-2">
          <SidebarEntry name="Events" href="/admin/dashboard/events" />
          <SidebarEntry name="Profil" href="/admin/dashboard/profil" />
          <SidebarEntry name="Indstillinger" href="/admin/dashboard/indstillinger" />
        </ul>
      </CardContent>
    </Card>
  )
}
