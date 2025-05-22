import { createContext, useContext, useEffect, useRef, type DialogHTMLAttributes, type HTMLAttributes, type PropsWithChildren } from "react";
import { FaXmark } from "react-icons/fa6";
import { cn } from "@/lib/clsx";

type ModalContext = {
	onClose: () => void;
}

const ModalContext = createContext<ModalContext | undefined>(undefined)

export const useModalContext = () => {
	const ctx = useContext(ModalContext)
	if (!ctx) throw new Error("No ModalContext.Provider found")

	return ctx
}

type Props = Omit<DialogHTMLAttributes<HTMLDialogElement>, "onClose"> & {
	show: boolean
	onClose: () => void;
}

const Modal = ({ show, onClose, children, className, ...rest }: Props) => {
	const ref = useRef<HTMLDialogElement>(null)

	useEffect(() => {
		if (!ref.current) return
		show ? ref.current.showModal() : ref.current.close();
	}, [show]);

	return (
		<ModalContext.Provider value={{ onClose }}>
			<dialog
				{...rest}
				ref={ref}
				onClose={onClose}
				className={cn(
					'fixed top-1/2 left-1/2 w-full min-w-xs -translate-x-1/2 -translate-y-1/2 flex-col overflow-hidden rounded-md border border-zinc-800 sm:min-w-lg',
					className
				)}
			>
				{children}
			</dialog>
		</ModalContext.Provider>
	)
}



const Header = ({ children }: PropsWithChildren) => {
	const { onClose } = useModalContext()

	return (
		<div className="text-text relative flex flex-col justify-center gap-2 bg-zinc-950 p-6" >
			{children}

			< button
				type="button"
				onClick={onClose}
				className="text-text/50 hover:text-text absolute top-6 right-6"
			>
				<FaXmark />
			</button >
		</div>
	)
}

const Title = ({ children, ...rest }: HTMLAttributes<HTMLHeadingElement>) => (
	<h1 className="font-medium" {...rest}>{children}</h1>
)

const Description = ({ children, ...rest }: HTMLAttributes<HTMLParagraphElement>) => (
	<p className="text-text/50 text-sm" {...rest}>{children}</p>
)

const Content = ({ children, className }: HTMLAttributes<HTMLDivElement>) => (
	<div className={cn('text-text max-h-64 w-full overflow-y-scroll bg-zinc-950 p-6', className)}>
		{children}
	</div>
)

const Footer = ({ children, className }: HTMLAttributes<HTMLDivElement>) => (
	<div className={cn('flex justify-end border-t border-t-zinc-800 bg-zinc-950 p-6', className)}>
		{children}
	</div>

)

Modal.Header = Header
Modal.Title = Title

Modal.Content = Content
Modal.Description = Description

Modal.Footer = Footer

export default Modal
