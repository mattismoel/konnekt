<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLLIElement> & {
		disabled: boolean;
	} & ({ action?: undefined; href: string } | { action: () => void; href?: undefined });

	let { action, href, disabled, children, ...rest }: Props = $props();
</script>

<li class={cn('text-text/85 hover:text-text py-2', rest.class)}>
	<svelte:element
		this={href ? 'a' : 'button'}
		type={href ? undefined : 'button'}
		href={href && !disabled ? href : undefined}
		aria-disabled={href ? disabled : undefined}
		role={href && disabled ? 'link' : undefined}
		tabindex={href && disabled ? -1 : 0}
		class="disabled:text-text/50 w-full text-left"
	>
		{@render children?.()}
	</svelte:element>
</li>
