<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLDivElement> = $props();
</script>

<div class={cn('max-h-64 w-full overflow-y-scroll bg-zinc-900 p-4', rest.class)}>
	{@render children?.()}
</div>
