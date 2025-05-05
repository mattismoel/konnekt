<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Position = { x: number; y: number };
	let mousePos = $state<Position>({ x: 0, y: 0 });

	type Props = HTMLAttributes<HTMLDivElement> & {};

	let { ...rest }: Props = $props();
</script>

<svelte:window
	onmousemove={(e) =>
		(mousePos = {
			x: e.clientX,
			y: e.clientY + e.currentTarget.scrollY
		})}
/>

<div
	style:top={`${mousePos.y}px`}
	style:left={`${mousePos.x}px`}
	class={cn(
		'pointer-events-none absolute z-10 hidden h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-white mix-blend-soft-light blur-[265px] brightness-75 sm:block',
		rest.class
	)}
></div>
