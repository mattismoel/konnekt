import { Link } from "@tanstack/react-router"

const Footer = () => (
	<footer className='flex flex-col gap-8 px-responsive py-12 border-t border-t-zinc-900'>
		<div className="grid grid-cols-1 sm:grid-cols-2 gap-8">
			<section>
				<h3 className="font-heading text-heading font-semibold mb-4">Find rundt</h3>
				<ul>
					<li><Link to="/">Hjem</Link></li>
					<li><Link to="/events">Events</Link></li>
					<li><Link to="/artists">Kunstnere</Link></li>
					<li><Link to="/about-ut">Om os</Link></li>
				</ul>
			</section>

			<section className="flex flex-col sm:items-end">
				<span className="font-bold mb-4 text-heading">Logo</span>
				<ul className="flex flex-col sm:items-end">
					<li><a href="mailto:konnekt.samarbejde@gmail.com">konnekt.samarbejde@gmail.com</a></li>
					<li><a href="mailto:booking.konnekt@gmail.com">booking.konnekt@gmail.com</a></li>
				</ul>
			</section>
		</div>

		<p className="text-center text-sm">&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen Konnekt</p>
	</footer>
)

export default Footer
