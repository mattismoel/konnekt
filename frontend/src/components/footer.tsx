export const Footer = () => (
  <footer className="px-auto border-t border-t-zinc-900 bg-zinc-950 pt-8 pb-6">
    <div className="mb-8 grid w-full grid-cols-1 sm:grid-cols-2">
      {/* NAVIGATION */}
      <div className="flex-1">
        <h3 className="mb-2 font-bold">Find rundt.</h3>
        <ul className="text-text/50">
          <li><a href="/">Hjem</a></li>
          <li><a href="/events">Events</a></li>
          <li><a href="/artists">Kunstnere</a></li>
          <li><a href="/om-os">Om os</a></li>
          <li><a href="/sponsorer">Sponsorer</a></li>
        </ul>
      </div>
      {/* CONTACT INFORMATION */}
      <div className="flex flex-1 flex-col items-end">
        <h3 className="mb-2 font-black">KONNEKT&reg;</h3>
        <ul className="text-text/50 flex flex-col items-end">
          <a href="tel:+4512345678">+45 12 34 56 78</a>
          <a href="mailto:konnekt@mail.dk">konnekt@mail.dk</a>
          <a href="https://google.maps.com/konnekt">Konnektvej 17B, Konnekticut</a>
        </ul>
      </div>
    </div>
    <div className="text-text/50 flex justify-center text-xs">
      <span>&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen&nbsp;KONNEKT&reg;</span>
    </div>
  </footer>
)
