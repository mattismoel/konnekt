<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLHeadingElement>;

	let { children, ...rest }: Props = $props();
</script>

<h1 {...rest} class={cn('text-xl font-bold', rest.class)}>{@render children?.()}</h1>
