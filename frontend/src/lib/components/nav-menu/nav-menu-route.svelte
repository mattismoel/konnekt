<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '$lib/clsx';
	import type { HTMLAnchorAttributes } from 'svelte/elements';

	type Props = HTMLAnchorAttributes;

	let { href, children, ...rest }: Props = $props();

	let isActive = $derived(page.url.pathname === href);

	$inspect(isActive, page.url.pathname);
</script>

<li>
	<a
		{...rest}
		{href}
		class={cn('text-text/50 text-3xl', {
			'text-text font-bold': isActive
		})}>{@render children?.()}</a
	>
</li>
