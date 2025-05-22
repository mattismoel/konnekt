import { cn } from '@/lib/clsx';
import { type HTMLAttributes } from 'react';
import { useMousePos } from '../hooks/useMousePos';

const GlowCursor = ({ className }: HTMLAttributes<HTMLDivElement>) => {
	const { x: mousePosX, y: mousePosY } = useMousePos()

	return (
		<div
			style={{
				left: `${mousePosX}px`,
				top: `${mousePosY}px`,
			}}
			className={cn(
				'pointer-events-none absolute z-10 hidden h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-white mix-blend-soft-light blur-[265px] brightness-150 sm:block',
				className
			)}
		></div>
	)
}

export default GlowCursor
