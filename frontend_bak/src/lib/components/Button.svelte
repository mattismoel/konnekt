<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = HTMLButtonAttributes & {
		expandX?: boolean;
		expandY?: boolean;
		variant?: 'primary' | 'secondary' | 'dangerous' | 'ghost';
	};

	let {
		children,
		class: className,
		expandX,
		expandY,
		variant = 'primary',
		...rest
	}: Props = $props();
</script>

<button
	type="button"
	class={cn(
		`flex max-w-64 items-center justify-center gap-1 rounded-sm px-4 py-2 font-medium text-zinc-950`,
		{
			'w-full max-w-none': expandX,
			'h-full': expandY,
			'bg-zinc-100 text-zinc-950 hover:bg-zinc-300': variant === 'primary',
			'border border-zinc-100 text-zinc-100 transition-colors hover:bg-zinc-300 hover:text-zinc-950':
				variant === 'secondary',
			'border border-red-700 bg-red-800 text-red-300 hover:text-red-100': variant === 'dangerous',
			'border border-zinc-900 font-normal text-zinc-500 hover:bg-zinc-900': variant === 'ghost'
		},
		className
	)}
	{...rest}
>
	{#if children}
		{@render children()}
	{/if}
</button>
