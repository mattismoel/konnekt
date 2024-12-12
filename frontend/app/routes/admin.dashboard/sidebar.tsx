import { Card, CardContent, CardFooter, CardHeader } from "@/components/ui/card"
import { SidebarEntry } from "./sidebar-entry"
import { useAuth } from "@/lib/context/auth.provider"
import { Button } from "@/components/ui/button"
import { cn } from "@/lib/utils"

type Props = {
  className?: string;
}

export const Sidebar = ({ className }: Props) => {
  const { logOut } = useAuth()

  return (
    <Card className={cn("", className)}>
      <CardHeader>
        <h3 className="font-black text-xl">KONNEKT &reg;</h3>
      </CardHeader>
      <CardContent>
        <ul className="flex flex-col gap-2">
          <SidebarEntry name="Events" href="/admin/dashboard/events" />
          <SidebarEntry name="Profil" href="/admin/dashboard/profil" />
          <SidebarEntry name="Indstillinger" href="/admin/dashboard/indstillinger" />
        </ul>
      </CardContent>
      <CardFooter>
        <Button variant="outline" onClick={logOut}>Log ud</Button>
      </CardFooter>
    </Card>
  )
}
