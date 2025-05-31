import Logo from '@/lib/assets/logo';
import { Link } from '@tanstack/react-router';
import { FaFacebook, FaInstagram, FaTiktok } from 'react-icons/fa6';
import type { IconType } from 'react-icons/lib';

type Src = {
	icon: IconType;
	title: string;
}

const socialMap = new Map<string, Src>([
	["https://www.instagram.com/konnekt_odense/", { icon: FaInstagram, title: "Instagram" }],
	["https://www.tiktok.com/@konnekt_", { icon: FaTiktok, title: "TikTok" }],
	["https://www.facebook.com/profile.php?id=61574860865073", { icon: FaFacebook, title: "Facebook" }],
])

const Footer = () => {
	return (
		<footer className="px-auto border-t border-t-zinc-900 bg-zinc-950 pt-8 pb-6">
			<div className="mb-8 grid w-full grid-cols-1 gap-8 sm:grid-cols-2">
				{/*  NAVIGATION  */}
				<div className="flex-1">
					<span className="font-heading mb-2 inline-block align-top leading-none font-bold">Find rundt</span>
					<ul className="text-text/50">
						<li><Link to="/">Hjem</Link></li>
						<li><Link to="/events">Events</Link></li>
						<li><Link to="/artists">Kunstnere</Link></li>
						<li><Link to="/about">Om os</Link></li>
					</ul>
				</div>
				{/* CONTACT INFORMATION  */}
				<div className="flex flex-1 flex-col items-start sm:items-end gap-4">
					<Logo className="mb-2 hidden h-4 sm:block" />
					<span className="font-heading mb-2 font-bold sm:hidden">Kontakt os</span>
					<address className="not-italic">
						<div className="text-text/50 flex flex-col items-start sm:items-end gap-2">
							<a href="mailto:konnekt.samarbejde@gmail.dk">konnekt.samarbejde@gmail.com</a>
							<a href="mailto:booking.konnekt@gmail.dk">booking.konnekt@gmail.com</a>
						</div>
					</address>
					<SocialMediaList socialMap={socialMap} />
				</div>
			</div >
			<div className="text-text/50 flex justify-center text-xs">
				<span>&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen&nbsp;Konnekt</span>
			</div>
		</footer >
	)
}

type SocialMediaListProps = {
	socialMap: Map<string, Src>
}

const SocialMediaList = ({ socialMap }: SocialMediaListProps) => (
	<ul className="flex gap-4 items-center text-xl text-text/50">
		{Array.from(socialMap).map(([href, { icon: Icon, title }]) => (
			<li key={href} className="hover:text-text">
				<a title={title} href={href}><Icon /></a>
			</li>
		))}
	</ul>
)

export default Footer
