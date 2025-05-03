<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLLIElement>;

	let { children, ...rest }: Props = $props();
</script>

<li
	{...rest}
	class={cn(
		'relative flex items-center justify-between gap-4 rounded-sm border border-zinc-900 p-3 hover:border-zinc-800 hover:bg-zinc-900 sm:border-transparent',
		rest.class
	)}
>
	{@render children?.()}
</li>
