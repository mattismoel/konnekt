import { cn } from "@/lib/clsx";
import { useEffect, useState, type HTMLAttributes } from "react";

const HIDE_TIMEOUT_DURATION_MS = 3000;

type Status = 'approved' | 'non-approved';

type Props = HTMLAttributes<HTMLDivElement> & {
	status: Status;
};

const MemberStatusIndicator = ({ status, className }: Props) => {
	let [visible, setVisible] = useState(true);

	useEffect(() => {
		if (status === 'non-approved') return;
		const timeout = setTimeout(() => {
			visible = false;
		}, HIDE_TIMEOUT_DURATION_MS);

		return () => {
			clearTimeout(timeout);
		};
	}, []);

	return (
		<div
			role="complementary"
			onFocus={() => setVisible(true)}
			onBlur={() => setVisible(false)}
			onMouseOver={() => setVisible(true)}
			onMouseLeave={() => setVisible(false)}
			// class:visible
			className={cn(
				'group flex h-6 w-6 cursor-default items-center justify-center overflow-hidden rounded-full border border-zinc-800 bg-zinc-900 px-1 transition-[width] ',
				{ "w-24 justify-between": visible },
				className
			)}
		>
			<div
				className={cn('h-3 w-3 rounded-full border border-green-400 bg-green-500', {
					'border-red-400 bg-red-500': status === 'non-approved'
				})}
			></div>
			<span className="hidden w-full text-center text-xs group-[.visible]:block">
				{status === 'approved'
					? "Godkendt"
					: status === "non-approved"
						? "Ikke-godkendt"
						: undefined}
			</span>
		</div>
	)
}

export default MemberStatusIndicator
