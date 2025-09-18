import { useScroll } from "@/lib/hooks/useScroll";
import { cn } from "@/lib/clsx";
import { Link, useRouterState } from "@tanstack/react-router";
import { FaBars } from "react-icons/fa6"
import Logo from "../logo";

type Props = {
	entries: { pathname: string, title: string }[]
	menuOpen: boolean;
	onOpenMenu: () => void;
}

const Navbar = ({ entries: links, menuOpen, onOpenMenu }: Props) => {
	const { location } = useRouterState()
	const { direction: { y: scrollDirY } } = useScroll()

	return (
		<nav className={cn("z-50 fixed w-full px-responsive h-nav flex justify-between items-center bg-gradient-to-b from-background/75 outline outline-transparent transition-colors [.scrolled]:to-background [.scrolled]:from-background [.scrolled]:outline-zinc-900", scrollDirY > 0 && "scrolled")}>
			<div className="flex gap-8 items-center">
				<button type="button" onClick={onOpenMenu} className={cn("transition-[rotate] md:hidden [.menu-open]:-rotate-90", menuOpen && "menu-open")}>
					<FaBars className="text-heading" />
				</button>
				<Link to="/" className="font-bold text-heading">
					<Logo className="h-4" />
				</Link>
			</div>
			<ul className="hidden text-heading/75 gap-8 md:flex">
				{links.map(({ pathname, title }) => (
					<NavEntry active={location.pathname === pathname} pathname={pathname} title={title} />
				))}
			</ul>
		</nav>
	)
}

type NavEntryProps = {
	pathname: string,
	title: string,
	active: boolean;
}

const NavEntry = ({ pathname, title, active }: NavEntryProps) => {
	return (
		<li>
			<Link title={title} to={pathname} className={cn("before:block before:content-[attr(title)] before:font-bold before:h-0 before:overflow-hidden before:invisible [.active]:font-bold", active && "active")}>
				{title}
			</Link>
		</li>
	)
}

export default Navbar
