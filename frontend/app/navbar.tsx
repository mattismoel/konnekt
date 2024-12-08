export const Navbar = () => {
  return (
    <nav className="h-nav flex items-center px-8 border-b border-gray-950 justify-between">
      <a href="/" className="font-black">KONNEKT&reg;</a>
      <ul className="flex gap-4">
        <li><a href="/events">Events</a></li>
        <li><a href="/om-os">Om os</a></li>
        <li><a href="/sponsorer">Sponsorer</a></li>
      </ul>
    </nav>
  )
}
