import { cn } from '@/lib/clsx';
import Fader from './fader';
import { useRef, useState, type PropsWithChildren } from 'react';

const Caroussel = ({ children }: PropsWithChildren) => {
	let [scrollX, setScroll] = useState(0);
	let inner = useRef<HTMLDivElement>(null);
	let outer = useRef<HTMLDivElement>(null);

	let [isScrolled, setIsScrolled] = useState(false);

	const updateScroll = (newScroll: number) => {
		if (!inner.current || !outer.current) return

		const innerWidth = inner.current?.scrollWidth;
		const outerWidth = outer.current?.getBoundingClientRect().width;

		const diff = innerWidth - outerWidth;
		setScroll(newScroll);

		setIsScrolled(scrollX >= diff);
	};
	return (
		<div className="relative isolate" ref={outer}>
			{/* RIGHT FADER */}
			<Fader
				direction="left"
				className={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
					'from-transparent': isScrolled
				})}
			/>
			<Fader
				direction="right"
				className={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
					'from-transparent': scrollX <= 0
				})}
			/>
			<div
				ref={inner}
				className="flex w-full gap-4 overflow-x-scroll"
				onScroll={(e) => updateScroll(e.currentTarget.scrollLeft)}
			>
				{children}
			</div>
		</div>
	)
}

export default Caroussel
