<script lang="ts">
	import { cn } from '$lib/clsx';
	import { clickOutside } from '$lib/hooks/click-outside.svelte';
	import type { HTMLAttributes } from 'svelte/elements';
	import CaretIcon from '~icons/mdi/caret';

	type Props = HTMLAttributes<HTMLElement> & {
		show: boolean;
	};

	let { show = $bindable(), children, ...rest }: Props = $props();
</script>

<div
	class:pointer-events-none={!show}
	class:opacity-50={show}
	class="fixed z-50 h-svh w-full bg-black opacity-0 transition-opacity"
></div>

<aside
	{...rest}
	use:clickOutside
	onclickoutside={() => (show = false)}
	class={cn(
		'fixed bottom-0 z-50 block w-screen translate-y-full px-2  transition-transform duration-300 ease-in-out',
		{
			'translate-y-0': show
		}
	)}
>
	<div
		class="flex h-full w-full flex-col justify-end rounded-t-md border border-zinc-800 bg-zinc-950 px-8 pt-12 pb-32"
	>
		{@render children?.()}
		<button
			type="button"
			onclick={() => (show = false)}
			class="absolute bottom-8 left-1/2 -translate-x-1/2 text-2xl"
			><CaretIcon
				class={cn('rotate-0 transition-transform duration-400', {
					'rotate-180': show
				})}
			/></button
		>
	</div>
</aside>
