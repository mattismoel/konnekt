import { Link } from "@tanstack/react-router"
import SocialList from "../social-list"
import Logo from "../logo"

type Props = {
	routes: { title: string, pathname: string }[]
	socials: string[]
	mails: string[]
}

const Footer = ({ routes, socials, mails }: Props) => (
	<footer className='flex flex-col gap-8 px-responsive py-12 border-t border-t-zinc-900'>
		<div className="grid grid-cols-1 sm:grid-cols-2 gap-8">
			<section>
				<h3 className="font-heading text-heading font-semibold mb-4">Find rundt</h3>
				<ul>
					{routes.map(({ title, pathname }) => (
						<li key={pathname}>
							<Link to={pathname} className="hover:underline">{title}</Link>
						</li>
					))}
				</ul>
			</section>

			<section className="flex gap-4 flex-col sm:items-end">
				<div className="h-full">
					<Logo className="h-4 text-heading" />
				</div>
				<ul className="flex flex-col sm:items-end">
					{mails.map(mail => (
						<li key={mail}>
							<a href={`mailto:${mail}`} className="hover:underline">{mail}</a>
						</li>
					))}
				</ul>
				<SocialList size="sm" urls={socials} />
			</section>
		</div>

		<p className="text-center text-sm">&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen Konnekt</p>
	</footer>
)

export default Footer
