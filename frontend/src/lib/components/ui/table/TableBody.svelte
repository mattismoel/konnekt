<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLDivElement> = $props();
</script>

<tbody class={cn('divide-y divide-zinc-900', rest.class)}>
	{@render children?.()}
</tbody>
