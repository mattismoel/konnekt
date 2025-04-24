<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	const HIDE_TIMEOUT_DURATION_MS = 3000;

	type Status = 'approved' | 'non-approved';

	type Props = HTMLAttributes<HTMLDivElement> & {
		status: Status;
	};

	let { status, ...rest }: Props = $props();

	let visible = $state(true);

	$effect(() => {
		if (status === 'non-approved') return;
		const timeout = setTimeout(() => {
			visible = false;
		}, HIDE_TIMEOUT_DURATION_MS);

		return () => {
			clearTimeout(timeout);
		};
	});
</script>

<div
	role="complementary"
	onfocus={() => (visible = true)}
	onblur={() => (visible = false)}
	onmouseover={() => (visible = true)}
	onmouseleave={() => (visible = false)}
	class:visible
	class={cn(
		'group flex h-6 w-6 cursor-default items-center justify-center overflow-hidden rounded-full border border-zinc-800 bg-zinc-900 px-1 transition-[width] [.visible]:w-24 [.visible]:justify-between',
		rest.class
	)}
>
	<div
		class={cn('h-3 w-3 rounded-full border border-green-400 bg-green-500', {
			'border-red-400 bg-red-500': status === 'non-approved'
		})}
	></div>
	<span class="hidden w-full text-center text-xs group-[.visible]:block">
		{#if status === 'approved'}
			Godkendt
		{:else if status === 'non-approved'}
			Ikke-godkendt
		{/if}
	</span>
</div>
