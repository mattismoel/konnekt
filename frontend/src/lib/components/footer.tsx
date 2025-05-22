import Logo from '@/lib/assets/logo';
import { Link } from '@tanstack/react-router';

const Footer = () => {
	return (
		<footer className="px-auto border-t border-t-zinc-900 bg-zinc-950 pt-8 pb-6">
			<div className="mb-8 grid w-full grid-cols-1 gap-8 sm:grid-cols-2">
				{/*  NAVIGATION  */}
				<div className="flex-1">
					<h3 className="font-heading mb-2 inline-block align-top leading-none font-bold">Find rundt</h3>
					<ul className="text-text/50">
						<li><Link to="/">Hjem</Link></li>
						<li><Link to="/events">Events</Link></li>
						<li><Link to="/artists">Kunstnere</Link></li>
						<li><Link to="/about">Om os</Link></li>
					</ul>
				</div>
				{/* CONTACT INFORMATION  */}
				<div className="flex flex-1 flex-col items-start sm:items-end">
					<Logo className="mb-2 hidden h-4 sm:block" />
					<h3 className="font-heading mb-2 font-bold sm:hidden">Kontakt os</h3>
					<address className="not-italic">
						<ul className="text-text/50 flex flex-col items-start sm:items-end">
							<a href="mailto:konnekt.samarbejde@gmail.dk">konnekt.samarbejde@gmail.com</a>
							<a href="mailto:booking.konnekt@gmail.dk">booking.konnekt@gmail.com</a>
						</ul>
					</address>
				</div>
			</div>
			<div className="text-text/50 flex justify-center text-xs">
				<span>&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen&nbsp;KONNEKT&reg;</span>
			</div>
		</footer>
	)
}

export default Footer
