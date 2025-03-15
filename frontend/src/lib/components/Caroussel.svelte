<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';
	import Fader from './ui/Fader.svelte';

	type Props = HTMLAttributes<HTMLDivElement> & {
		title?: string;
	};

	let { children }: Props = $props();

	let scrollX = $state(0);
	let inner: HTMLDivElement;
	let outer: HTMLDivElement;

	let isScrolledToRight = $state(false);

	const updateScroll = (newScrollX: number) => {
		const innerWidth = inner.scrollWidth;
		const outerWidth = outer.getBoundingClientRect().width;

		const diff = innerWidth - outerWidth;
		scrollX = newScrollX;

		isScrolledToRight = scrollX >= diff;
	};
</script>

<div class="relative isolate" bind:this={outer}>
	<!--RIGHT FADER -->
	<Fader
		direction="left"
		class={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
			'from-transparent': isScrolledToRight
		})}
	/>
	<Fader
		direction="right"
		class={cn('absolute z-10 w-32 from-zinc-950 transition-colors duration-300', {
			'from-transparent': scrollX <= 0
		})}
	/>
	<div
		bind:this={inner}
		class="flex w-full gap-4 overflow-x-scroll"
		onscroll={(e) => updateScroll(e.currentTarget.scrollLeft)}
	>
		{@render children?.()}
	</div>
</div>
