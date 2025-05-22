import {
	forwardRef,
	useRef,
	type AnchorHTMLAttributes,
	type ButtonHTMLAttributes,
	type HTMLAttributes,
} from 'react';

import { useOnClickOutside } from '../hooks/useClickOutside';
import { cn } from '../clsx';
import { createLink } from '@tanstack/react-router';
import { FaBars, FaEllipsis } from 'react-icons/fa6';

type Props = HTMLAttributes<HTMLDivElement> & {
	show: boolean;
	onClose: () => void;
};

const ContextMenu = ({ show, children, className, onClose }: Props) => {
	const ref = useRef<HTMLDivElement>(null)

	useOnClickOutside(ref, onClose)

	return (
		<div
			ref={ref}
			role="menu"
			tabIndex={0}
			className={cn(
				'absolute top-1/2 right-4 z-50 hidden min-w-48 flex-col divide-y divide-zinc-900 overflow-hidden rounded-md border border-zinc-900 bg-zinc-950', {
				"flex": show
			}, className)}
		>
			{children}
		</div>
	)
}


type LinkEntryProps = AnchorHTMLAttributes<HTMLAnchorElement> & {
	disabled?: boolean;
}

const entryBaseClasses = "disabled:text-text/25 w-full px-4 py-2 text-left hover:bg-zinc-900"

const LinkEntry = createLink(
	forwardRef<HTMLAnchorElement, LinkEntryProps>(
		({ href, disabled, children, className, ...rest }, ref) => (
			<a
				ref={ref}
				href={href && !disabled ? href : undefined}
				aria-disabled={disabled}
				role={href && disabled ? "link" : undefined}
				tabIndex={href && disabled ? -1 : 0}
				className={cn(entryBaseClasses, className)}
				{...rest}
			>
				{children}
			</a>
		)
	)
)

type EntryProps = Omit<ButtonHTMLAttributes<HTMLButtonElement>, "onClick"> & {
	onClick: () => void;
}

const Entry = forwardRef<HTMLButtonElement, EntryProps>(({ children, className, ...rest }, ref) => (
	<button
		ref={ref}
		{...rest}
		className={cn(entryBaseClasses, className)}
		type="button"
	>
		{children}
	</button>
))

const Button = forwardRef<HTMLButtonElement, ButtonHTMLAttributes<HTMLButtonElement>>((props, ref) => (
	<button ref={ref} {...props} type="button" className="rounded-md p-1 hover:bg-zinc-800">
		<FaEllipsis />
	</button>
))

ContextMenu.Entry = Entry
ContextMenu.LinkEntry = LinkEntry
ContextMenu.Button = Button

export default ContextMenu
