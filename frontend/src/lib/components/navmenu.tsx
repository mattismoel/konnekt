import { cn } from '@/lib/clsx';
import { forwardRef, useRef, type AnchorHTMLAttributes, type ButtonHTMLAttributes, type HTMLAttributes } from 'react';
import { BiCaretUp } from 'react-icons/bi';
import { useOnClickOutside } from '../hooks/useClickOutside';
import { createLink, useLocation } from '@tanstack/react-router';

type Props = HTMLAttributes<HTMLElement> & {
	expanded: boolean;
	onClose: () => void;
};

type OverlayProps = {
	expanded: boolean;
}

const BackgroundOverlay = ({ expanded }: OverlayProps) => (
	<div
		className={cn("fixed z-50 h-svh w-full bg-black opacity-0 transition-opacity", {
			"pointer-events-none": !expanded,
			"opacity-50": expanded
		})}
	/>
)

type CloseButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
	expanded: boolean;
}

const CloseButton = forwardRef<HTMLButtonElement, CloseButtonProps>(({ expanded, ...rest }, ref) => {
	return (
		<button
			ref={ref}
			{...rest}
			type="button"
			className="absolute bottom-8 left-1/2 -translate-x-1/2 text-2xl"
		>
			<BiCaretUp
				className={cn('rotate-0 transition-transform duration-400', { 'rotate-180': expanded })}
			/>
		</button
		>
	)
})

const NavMenu = ({ expanded, children, onClose, ...rest }: Props) => {
	const ref = useRef<HTMLElement>(null)
	useOnClickOutside(ref, onClose)

	return (
		<>
			<BackgroundOverlay expanded={expanded} />

			<aside
				ref={ref}
				{...rest}
				className={cn(
					'fixed bottom-0 z-50 block w-screen translate-y-full px-2  transition-transform duration-300 ease-in-out',
					{ 'translate-y-0': expanded }
				)}
			>
				<div className="flex h-full w-full flex-col justify-end rounded-t-md border border-zinc-800 bg-zinc-950 px-8 pt-12 pb-32">
					{children}
					<CloseButton expanded={expanded} onClick={onClose} />
				</div>
			</aside>
		</>
	)
}

type RouteListProps = HTMLAttributes<HTMLUListElement>

const RouteList = ({ children, ...rest }: RouteListProps) => {
	return (
		<ul {...rest} className="flex flex-col gap-4">
			{children}
		</ul>
	)
}

type RouteEntryProps = AnchorHTMLAttributes<HTMLAnchorElement>

const RouteEntry = createLink(
	forwardRef<HTMLAnchorElement, RouteEntryProps>(
		({ href, children, className, ...rest }, ref) => {
			const { pathname } = useLocation()

			const isActive = pathname === href;

			return (
				<li className="cursor-pointer">
					<a
						ref={ref}
						{...rest}
						className={cn("text-text/50 text-3xl",
							{ 'text-text font-bold': isActive },
						)}
					>
						{children}
					</a>
				</li>
			)
		}
	)
)

NavMenu.RouteList = RouteList
NavMenu.RouteEntry = RouteEntry

export default NavMenu;
