import { forwardRef, type AnchorHTMLAttributes, type HTMLAttributes, type LiHTMLAttributes } from "react";
import { cn } from "../../clsx";
import { createLink } from "@tanstack/react-router";

const List = ({ children, ...rest }: HTMLAttributes<HTMLUListElement>) => (
	<ul {...rest} className="space-y-2">
		{children}
	</ul>
)

const Entry = ({ className, ...rest }: LiHTMLAttributes<HTMLLIElement>) => (
	<li
		{...rest}
		className={cn(
			'relative flex items-center justify-between gap-4 rounded-sm border border-zinc-900 hover:border-zinc-800 hover:bg-zinc-900 sm:border-transparent',
			className
		)}
	/>
)

const sectionClasses = (disabled: boolean, className?: string) => cn(
	"w-full flex flex-col p-3 cursor-pointer", className,
	{ "opacity-50 cursor-default": disabled },
)

type SectionBaseProps = {
	disabled?: boolean;
}

type SectionProps = SectionBaseProps & HTMLAttributes<HTMLDivElement>

const Section = ({ disabled = false, children, className, ...rest }: SectionProps) => (
	<div {...rest} className={sectionClasses(disabled, className)}>
		{children}
	</div>
)

type LinkEntryProps = SectionBaseProps & AnchorHTMLAttributes<HTMLAnchorElement>

const LinkSection = createLink(
	forwardRef<HTMLAnchorElement, LinkEntryProps>(
		({ href, disabled = false, className, ...rest }, ref) => (
			<a ref={ref} {...rest} className={sectionClasses(disabled, className)} />
		)
	)
)

List.Entry = Entry
Entry.Section = Section
Entry.LinkSection = LinkSection

export default List
