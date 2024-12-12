import { cn } from "@/lib/utils";
import { BiCalendarAlt, BiCog, BiLogOut, BiSolidUserCircle } from "react-icons/bi";

type Props = {
  className?: string;
}

export const BottomBar = ({ className }: Props) => {
  return (
    <div
      className={cn(
        "bg-background border-t left-0 py-3 w-full flex items-center justify-between",
        "px-auto",
        className,
      )}
    >
      <a href="/admin/dashboard/events" className="flex flex-col items-center">
        <BiCalendarAlt className="text-2xl" />
        Events
      </a>
      <a href="/admin/dashboard/indstillinger" className="flex flex-col items-center">
        <BiCog className="text-2xl" />
        Indstillinger
      </a>
      <a href="/admin/dashboard/events" className="flex flex-col items-center">
        <BiSolidUserCircle className="text-2xl" />
        Profil
      </a>
      <a href="/admin/dashboard/events" className="flex flex-col items-center">
        <BiLogOut className="text-2xl" />
        Log ud
      </a>
    </div>
  )
}
