import type { HTMLAttributes } from "react";

type Props = HTMLAttributes<HTMLDivElement> & {
	srcs: Map<string, string>;
};

const LogoScroller = ({ srcs }: Props) => {
	return (
		<div className="group flex gap-8 overflow-hidden">
			{[...Array(2).keys()].map(idx => (
				<Track srcs={srcs} idx={idx} key={idx} />
			))}
		</div>
	)
}

type TrackProps = {
	srcs: Map<string, string>
	idx: number;
}

const Track = ({ srcs, idx }: TrackProps) => {
	return (
		<div className="animate-slide flex h-8 space-x-8" aria-hidden={idx > 0}>
			{Array.from(srcs).map(([name, src]) => (
				<img key={src} src={src} alt={name} className="max-w-none" />
			))}
		</div>
	)
}

export default LogoScroller
