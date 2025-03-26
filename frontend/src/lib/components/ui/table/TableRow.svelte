<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLTableRowElement> = $props();
</script>

<tr class={cn('transition-colors hover:bg-zinc-900', rest.class)}>
	{@render children?.()}
</tr>
