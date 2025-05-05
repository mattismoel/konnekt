<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes, HTMLAnchorAttributes } from 'svelte/elements';

	type Props = ((HTMLAttributes<HTMLDivElement> & { href?: undefined }) | HTMLAnchorAttributes) & {
		expand?: boolean;
	};

	let { expand = true, href, children, ...rest }: Props = $props();
</script>

<svelte:element
	this={href ? 'a' : 'div'}
	{...rest}
	href={href ? href : undefined}
	class:w-full={expand}
	class={cn('flex flex-col p-3', rest.class)}
>
	{@render children?.()}
</svelte:element>
