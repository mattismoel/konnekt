import { useLocation } from "@remix-run/react"
import { cn } from "@/lib/utils"

export const SidebarEntry = ({ name, href }: { name: string, href: string }) => {
  const { pathname } = useLocation()

  return (
    <a
      href={href}
      className={cn(
        "text-neutral-500",
        href === pathname && "font-bold text-neutral-100",
      )}
    >
      {name}
    </a>
  )
}
