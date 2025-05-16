import { cn } from '@/lib/clsx';
import { Link, useLocation } from "@tanstack/react-router";
import { useScroll } from '@/lib/hooks/useScroll';
import type { HTMLAttributes, PropsWithChildren } from 'react';

const Navbar = ({ children }: PropsWithChildren) => {
	const { y: scrollY } = useScroll()
	return (
		<nav
			className={cn("h-nav px-auto fixed z-40 flex w-screen items-center justify-between bg-gradient-to-b from-black/80 outline outline-transparent transition-colors duration-500 [.scrolled]:from-zinc-950 [.scrolled]:to-zinc-950 [.scrolled]:outline-zinc-800", {
				"scrolled": scrollY > 0
			})}
		>
			{children}
		</nav>
	)
}

const Header = ({ children }: PropsWithChildren) => {
	return (
		<div className="flex items-center gap-6">
			{children}
		</div>
	)
}

const Content = ({ children }: PropsWithChildren) => {
	return (
		<div className="flex items-center gap-8">
			{children}
		</div>
	)
}

type RouteEntryProps = {
	pathname: string;
	name: string;
};

const RouteList = ({ children, className }: HTMLAttributes<HTMLUListElement>) => {
	return (
		<ul className={cn('hidden items-center gap-6 text-lg text-zinc-50 md:flex', className)}>
			{children}
		</ul>
	)
}

const RouteEntry = ({ pathname, name }: RouteEntryProps) => {
	const { pathname: pagePathname } = useLocation()

	return (
		<li>
			<Link
				to={pathname}
				title={name}
				className={cn("hover:text-text text-text/75 [.is-current]:text-text transition-colors before:invisible before:block before:h-0 before:overflow-hidden before:font-medium before:content-[attr(title)] [.is-current]:font-medium", {
					"is-current": pathname === pagePathname
				})}
			>
				{name}
			</Link>
		</li>
	)
}

Navbar.Content = Content
Navbar.Header = Header
Navbar.RouteEntry = RouteEntry
Navbar.RouteList = RouteList

export default Navbar
