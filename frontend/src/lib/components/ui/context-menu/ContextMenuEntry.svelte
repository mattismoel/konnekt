<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

	type Props = {
		disabled: boolean;
	} & (
		| (Omit<HTMLAnchorAttributes, 'href'> & {
				href: string;
				onclick?: undefined;
		  })
		| (Omit<HTMLButtonAttributes, 'onclick'> & {
				onclick: () => void;
				href?: undefined;
		  })
	);

	let { href, disabled, children, ...rest }: Props = $props();
</script>

<svelte:element
	this={href ? 'a' : 'button'}
	type={href ? undefined : 'button'}
	href={href && !disabled ? href : undefined}
	aria-disabled={href ? disabled : undefined}
	role={href && disabled ? 'link' : undefined}
	tabindex={href && disabled ? -1 : 0}
	class={cn('disabled:text-text/25 w-full px-4 py-2 text-left hover:bg-zinc-900', rest.class)}
	{...rest}
>
	{@render children?.()}
</svelte:element>
