<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { Snippet } from 'svelte';

	type Props = {
		children: Snippet<[]>;
		expandX?: boolean;
		expandY?: boolean;
		extraClass?: string;
		type?: 'primary' | 'secondary' | 'dangerous';
	};

	let { children, expandX, expandY, extraClass, type = 'primary' }: Props = $props();
</script>

<button
	type="button"
	class={cn(
		`max-w-64 rounded-sm px-4 py-2 font-semibold text-zinc-950`,
		{
			'w-full max-w-none': expandX,
			'h-full': expandY,
			'bg-zinc-100 text-zinc-950 hover:bg-zinc-300': type === 'primary',
			'border border-zinc-100 text-zinc-100 transition-colors hover:bg-zinc-300 hover:text-zinc-950':
				type === 'secondary',
			'border border-red-700 bg-red-800 text-red-300 hover:text-red-100': type === 'dangerous'
		},
		extraClass
	)}
>
	{@render children()}
</button>
