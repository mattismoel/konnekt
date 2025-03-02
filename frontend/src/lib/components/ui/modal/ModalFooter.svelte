<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLDivElement> = $props();
</script>

<div class={cn('flex justify-end border-t border-t-zinc-800 bg-zinc-900 p-4', rest.class)}>
	{@render children?.()}
</div>
