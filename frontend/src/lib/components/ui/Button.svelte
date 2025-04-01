<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = HTMLButtonAttributes & {
		variant?: 'primary' | 'secondary' | 'dangerous' | 'ghost';
	};

	let { children, class: className, variant = 'primary', ...rest }: Props = $props();
</script>

<button
	type="button"
	class={cn(
		`flex h-min w-fit items-center justify-center gap-1 rounded-sm px-6 py-2 font-medium text-zinc-950`,
		{
			'bg-zinc-100 text-zinc-950 hover:bg-zinc-300': variant === 'primary',
			'border border-zinc-100 text-zinc-100 transition-colors hover:bg-zinc-100 hover:text-zinc-950':
				variant === 'secondary',
			'border border-red-900 bg-red-950 text-red-400 hover:border-red-800 hover:bg-red-900 hover:text-red-300':
				variant === 'dangerous',
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
