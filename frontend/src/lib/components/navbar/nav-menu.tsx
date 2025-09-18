import { cn } from "@/lib/clsx";
import { useClickOutside } from "@/lib/hooks/useClickOutside";
import { Link, useRouterState } from "@tanstack/react-router";
import { useRef } from "react";
import { FaChevronDown } from "react-icons/fa6";

type Props = {
	entries: {
		pathname: string;
		title: string;
	}[]
	show: boolean;
	onClose: () => void;
}

const NavMenu = ({ entries, show, onClose }: Props) => {
	const { location } = useRouterState()
	const ref = useRef<HTMLDivElement>(null)

	useClickOutside(ref, onClose)

	return (
		<div className={cn("group z-50 fixed top-0 left-0 w-screen h-screen [.show]:bg-black/75 pointer-events-none transition-colors [.show]:pointer-events-auto", show && "show")}>
			<div className="fixed w-full left-1/2 -translate-x-1/2 -bottom-full px-responsive transition-[bottom] group-[.show]:bottom-0">
				<div ref={ref} className="relative z-50 h-full bg-background p-8 border border-zinc-900 rounded-sm">
					<ul className="flex flex-col gap-4">
						{entries.map(({ pathname, title }) => (
							<NavEntry pathname={pathname} title={title} active={pathname === location.pathname} />
						))}
					</ul>
					<button type="button" onClick={onClose} className="w-full mt-16 flex justify-center rotate-180 transition-[rotate] group-[.show]:rotate-0"><FaChevronDown /></button>
				</div>
			</div>
		</div>
	)
}

type NavEntryProps = {
	pathname: string
	title: string
	active: boolean
}

const NavEntry = ({ pathname, title, active }: NavEntryProps) => (
	<li className={cn("text-3xl font-heading [.active]:font-bold [.active]:text-heading", active && "active")}>
		<Link to={pathname}>{title}</Link>
	</li>
)

export default NavMenu
