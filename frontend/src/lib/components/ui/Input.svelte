<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLInputAttributes } from 'svelte/elements';
	import FieldError from './FieldError.svelte';

	type Props = Omit<HTMLInputAttributes, 'class' | 'placeholder'> & {
		label: string;
		nonEmpty?: boolean;
		errors?: string[];
	};

	let { nonEmpty, errors, value = $bindable(''), label, ...rest }: Props = $props();

	let focused = $state(false);
</script>

<div class="space-y-2">
	<div class="relative">
		<label
			for={rest.name}
			class={cn(
				'text-normal pointer-events-none absolute top-1/2 left-4 -translate-y-1/2  transition-all duration-100',
				{
					'-top-3 left-0 text-xs text-zinc-300': focused || value !== '' || nonEmpty,
					'text-zinc-500': value === '' && !focused
				}
			)}>{label}</label
		>
		<input
			bind:value
			onfocus={() => (focused = true)}
			onblur={() => (focused = false)}
			class="bg-background rounded-sm border border-zinc-900 px-4 py-2"
			{...rest}
		/>
	</div>
	<FieldError {errors} />
</div>
