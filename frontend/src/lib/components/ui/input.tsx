import { cn } from '@/lib/clsx';
import { forwardRef, type InputHTMLAttributes } from 'react';

const Input = forwardRef<HTMLInputElement, InputHTMLAttributes<HTMLInputElement>>(({ className, ...rest }, ref) => {
	return (
		<input
			ref={ref}
			{...rest}
			className={cn("bg-background disabled:text-text/50 w-full rounded-sm border border-zinc-900 px-3 py-2", className)}
		/>
	)
})

export default Input
