<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		open: boolean;
		onClose: () => void;
	};

	let { open, onClose, children, ...rest }: Props = $props();

	let mouseHovering = $state(false);
</script>

<svelte:window onmousedown={() => (open && !mouseHovering ? onClose() : null)} />

<div
	role="menu"
	tabindex="0"
	onmouseenter={() => (mouseHovering = true)}
	onmouseleave={() => (mouseHovering = false)}
	class:open
	class={cn(
		'absolute top-1/2 right-4 z-50 hidden min-w-48 flex-col divide-y divide-zinc-900 overflow-hidden rounded-md border border-zinc-900 bg-zinc-950 [.open]:flex',
		rest.class
	)}
>
	{@render children?.()}
</div>
