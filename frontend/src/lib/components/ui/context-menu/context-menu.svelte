<script lang="ts">
	import { cn } from '$lib/clsx';
	import { clickOutside } from '$lib/hooks/click-outside.svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		show: boolean;
	};

	let { show = $bindable(false), children, ...rest }: Props = $props();
</script>

<div
	class:show
	use:clickOutside
	onclickoutside={() => (show = false)}
	role="menu"
	tabindex="0"
	class={cn(
		'absolute top-1/2 right-4 z-50 hidden min-w-48 flex-col divide-y divide-zinc-900 overflow-hidden rounded-md border border-zinc-900 bg-zinc-950 [.show]:flex',
		rest.class
	)}
>
	{@render children?.()}
</div>
