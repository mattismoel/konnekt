<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLLiAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLLiAttributes = $props();
</script>

<li
	{...rest}
	class={cn(
		'relative flex items-center justify-between gap-4 rounded-sm border border-zinc-900 hover:border-zinc-800 hover:bg-zinc-900 sm:border-transparent',
		rest.class
	)}
>
	{@render children?.()}
</li>
