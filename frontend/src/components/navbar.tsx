import { Link } from "@tanstack/react-router";

export const Navbar = () => (
  <nav
    className="h-nav absolute z-50 flex w-screen items-center justify-between bg-gradient-to-b from-black/65 px-8"
  >
    <a href="/" className="font-black">KONNEKT&reg;</a>
    <ul className="*:hover:text-text flex gap-6 text-zinc-50">
      <li><Link to="/events/">Events</Link></li>
      <li><Link to="/artists/">Kunstnere</Link></li>
      <li><Link to="/sponsorer/">Sponsorer</Link></li>
      <li><Link to="/om-os/">Om os</Link></li>
    </ul>
  </nav>
)
