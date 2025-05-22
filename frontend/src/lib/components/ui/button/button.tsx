import { type ButtonHTMLAttributes } from 'react';
import { baseClasses, variantClasses, type RootProps } from './base';
import { cn } from '@/lib/clsx';

type Props = RootProps & ButtonHTMLAttributes<HTMLButtonElement>

const Button = ({ children, type = "button", variant = "primary", className, ...rest }: Props) => {
	return (
		<button
			type={type}
			{...rest as Props}
			className={cn(baseClasses, variantClasses.get(variant), className)}
		>
			{children}
		</button>
	)
}

export default Button
