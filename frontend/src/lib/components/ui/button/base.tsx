import type { PropsWithChildren } from "react";

type Variant = 'outline' | 'primary' | 'secondary' | 'dangerous' | 'ghost';

export type RootProps = PropsWithChildren & {
	variant?: Variant;
	disabled?: boolean;
};

export const baseClasses =
	'flex h-min w-fit items-center justify-center gap-3 rounded-sm px-3 py-2 font-medium text-zinc-950 transition-colors disabled:opacity-50';

export const variantClasses = new Map<Variant, string>([
	['primary', 'bg-zinc-100 text-zinc-950 hover:bg-zinc-300'],
	[
		'secondary',
		'text-text/85 hover:text-text border border-zinc-700 bg-zinc-800 font-normal hover:border-zinc-600 hover:bg-zinc-700'
	],
	[
		'outline',
		'text-text/85 hover:text-text border border-zinc-100/75 transition-colors hover:border-zinc-100/30 hover:bg-zinc-100/25'
	],
	[
		'dangerous',
		'border border-red-900 bg-red-950 text-red-400 hover:border-red-800 hover:bg-red-900 hover:text-red-300'
	],
	['ghost', 'text-text/75 hover:text-text border border-zinc-900 font-normal hover:bg-zinc-900']
]);
