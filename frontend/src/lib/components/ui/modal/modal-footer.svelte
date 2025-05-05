<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLDivElement> = $props();
</script>

<div class={cn('flex justify-end border-t border-t-zinc-800 bg-zinc-950 p-6', rest.class)}>
	{@render children?.()}
</div>
