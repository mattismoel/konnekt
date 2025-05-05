<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLDivElement> = $props();
</script>

<div class={cn('rounded-md border border-zinc-800 bg-zinc-950', rest.class)}>
	{#if children}
		{@render children()}
	{/if}
</div>
