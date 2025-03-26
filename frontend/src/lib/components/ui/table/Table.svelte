<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLTableElement> = $props();
</script>

<table class={cn('w-full', rest.class)}>
	{@render children?.()}
</table>
