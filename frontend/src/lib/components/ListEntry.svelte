<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLLIElement>;

	let { children, ...rest }: Props = $props();
</script>

<li
	{...rest}
	class={cn(
		'relative flex items-center justify-between gap-4 rounded-sm border border-transparent p-3 hover:border-zinc-800 hover:bg-zinc-900',
		rest.class
	)}
>
	{@render children?.()}
</li>
