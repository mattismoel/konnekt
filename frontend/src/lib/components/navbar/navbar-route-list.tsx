import { cn } from '@/lib/clsx';
import type { HTMLAttributes } from 'react';

const routeList = ({ children, className }: HTMLAttributes<HTMLUListElement>) => {
	return (
		<ul className={cn('hidden items-center gap-6 text-lg text-zinc-50 md:flex', className)}>
			{children}
		</ul>
	)
}

export default routeList
