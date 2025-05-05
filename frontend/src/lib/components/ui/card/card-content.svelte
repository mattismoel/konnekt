<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement>;

	let { children, ...rest }: Props = $props();
</script>

<div {...rest} class={cn('flex flex-col p-6', rest.class)}>
	{@render children?.()}
</div>
