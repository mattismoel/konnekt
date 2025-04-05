<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLUListElement> & {
		open: boolean;
		onClose: () => void;
	};

	let { open, onClose, children, ...rest }: Props = $props();

	let mouseHovering = $state(false);
</script>

<svelte:window onmousedown={() => (open && !mouseHovering ? onClose() : null)} />

<ul
	onmouseenter={() => (mouseHovering = true)}
	onmouseleave={() => (mouseHovering = false)}
	class:open
	class={cn(
		'z-50 hidden min-w-48 divide-y divide-zinc-900 rounded-md border border-zinc-900 bg-zinc-950 px-4 py-2 [.open]:block',
		rest.class
	)}
>
	{@render children?.()}
</ul>
