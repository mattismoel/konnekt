<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLButtonAttributes } from 'svelte/elements';

	type Props = HTMLButtonAttributes & {
		variant?: 'outline' | 'primary' | 'secondary' | 'dangerous' | 'ghost';
	};

	let { children, variant = 'primary', ...rest }: Props = $props();
</script>

<button
	type="button"
	{...rest}
	class={cn(
		'flex h-min w-fit items-center justify-center gap-3 rounded-sm px-3 py-2 font-medium text-zinc-950 transition-colors disabled:opacity-50',
		{
			'bg-zinc-100 text-zinc-950 hover:bg-zinc-300': variant === 'primary',
			'text-text/85 hover:text-text border border-zinc-700 bg-zinc-800 font-normal hover:border-zinc-600 hover:bg-zinc-700':
				variant === 'secondary',
			'text-text/85 hover:text-text border border-zinc-100/75 transition-colors hover:border-zinc-100/30 hover:bg-zinc-100/25':
				variant === 'outline',
			'border border-red-900 bg-red-950 text-red-400 hover:border-red-800 hover:bg-red-900 hover:text-red-300':
				variant === 'dangerous',
			'border border-zinc-900 font-normal text-zinc-500 hover:bg-zinc-900': variant === 'ghost'
		},
		rest.class
	)}
>
	{#if children}
		{@render children()}
	{/if}
</button>
