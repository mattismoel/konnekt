import type { HTMLAttributes, PropsWithChildren } from "react";
import { cn } from "../clsx";

const AdminHeader = ({ children }: PropsWithChildren) => (
	<header className="flex flex-col gap-6">
		{children}
	</header>
)

const Title = ({ children, className, ...rest }: HTMLAttributes<HTMLHeadingElement>) => (
	<h1 {...rest} className={cn("font-heading text-4xl font-bold", className)}>
		{children}
	</h1>
)

const Description = ({ children, className, ...rest }: HTMLAttributes<HTMLParagraphElement>) => (
	<p {...rest} className={cn("text-text/50", className)}>
		{children}
	</p>
)

const Actions = ({ children }: PropsWithChildren) => (
	<div className="flex gap-4">{children}</div>
)

AdminHeader.Title = Title
AdminHeader.Description = Description
AdminHeader.Actions = Actions

export default AdminHeader
