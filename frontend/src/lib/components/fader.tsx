import { cn } from '@/lib/clsx';
import type { HTMLAttributes } from 'react';

type Direction = 'left' | 'right' | 'up' | 'down';

type Props = HTMLAttributes<HTMLDivElement> & {
	direction: Direction;
};

const Fader = ({ direction, className }: Props) => {
	return (
		<div
			className={cn(
				'pointer-events-none from-black',
				{
					'top-0 left-0 h-full min-w-4 bg-gradient-to-r': direction === 'right',
					'top-0 right-0 h-full min-w-4 bg-gradient-to-l': direction === 'left',
					'bottom-0 left-0 min-h-4 w-full bg-gradient-to-t': direction === 'up',
					'top-0 left-0 min-h-4 w-full bg-gradient-to-b': direction === 'down'
				},
				className
			)}
		></div>
	)
}

export default Fader
