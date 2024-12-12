import { cn } from "@/lib/utils"
import { useState } from "react"
import { BiMenu, BiX } from "react-icons/bi"

export const Navbar = () => {
  const [showPopover, setShowPopover] = useState(false)

  return (
    <>
      <nav className="bg-background h-nav flex items-center px-8 border-b border-gray-950 justify-between">
        <a href="/" className="font-black">KONNEKT&reg;</a>
        <ul className="gap-4 hidden sm:flex font-medium">
          <li><a href="/events">Events</a></li>
          <li><a href="/om-os">Om os</a></li>
          <li><a href="/sponsorer">Sponsorer</a></li>
        </ul>
        <button className="text-2xl sm:hidden" onClick={() => setShowPopover(true)}><BiMenu /></button>
      </nav>
      <div
        onClick={() => setShowPopover(false)}
        className={cn(
          "fixed opacity-95 top-0 left-0 h-screen w-screen bg-background z-50 flex items-center justify-center",
          "transition-all",
          { "opacity-0 pointer-events-none": showPopover === false },
        )}>
        <ul className="text-4xl space-y-4 font-semibold">
          <button className="fixed top-4 right-8 text-2xl shrink-0" onClick={() => setShowPopover(false)}>
            <BiX />
          </button>
          <li><a href="/events">Events.</a></li>
          <li><a href="/om-os">Om os.</a></li>
          <li><a href="/sponsorer">Sponsorer.</a></li>
        </ul>
      </div>
    </>
  )
}
