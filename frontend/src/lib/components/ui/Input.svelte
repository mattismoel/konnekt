<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLInputAttributes } from 'svelte/elements';

	type Props = Omit<HTMLInputAttributes, 'class' | 'placeholder'> & {
		label: string;
		expandX?: boolean;
		nonEmpty?: boolean;
	};

	let { expandX = true, nonEmpty, value = $bindable(''), label, ...rest }: Props = $props();

	let focused = $state(false);
</script>

<div class={cn('relative', { 'w-full flex-1': expandX })}>
	<label
		for={rest.name}
		class={cn('text-normal absolute top-1/2 left-4 -translate-y-1/2  transition-all duration-100', {
			'-top-3 left-0 text-xs text-zinc-300': focused || value !== '' || nonEmpty,
			'text-zinc-500': value === '' && !focused
		})}>{label}</label
	>
	<input
		{...rest}
		onfocus={() => (focused = true)}
		onblur={() => (focused = false)}
		class={cn('bg-background rounded-sm border border-zinc-900 px-4 py-2', {
			'w-full flex-1': expandX
		})}
		bind:value
	/>
</div>
