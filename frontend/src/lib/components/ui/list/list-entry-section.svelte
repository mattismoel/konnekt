<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes, HTMLAnchorAttributes } from 'svelte/elements';

	type Props = (
		| (HTMLAttributes<HTMLDivElement> & {
				href?: undefined;
		  })
		| (Omit<HTMLAnchorAttributes, 'href'> & {
				href: string;
		  })
	) & {
		expand?: boolean;
		disabled?: boolean;
	};

	let { disabled, expand = true, href, children, ...rest }: Props = $props();
</script>

<svelte:element
	this={href ? 'a' : 'div'}
	{...rest}
	href={href && !disabled ? href : undefined}
	class:w-full={expand}
	class:disabled
	class={cn('flex flex-col p-3 [.disabled]:cursor-default [.disabled]:opacity-50', rest.class)}
>
	{@render children?.()}
</svelte:element>
