<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLInputAttributes } from 'svelte/elements';
	import FieldError from './FieldError.svelte';

	type Props = HTMLInputAttributes & {
		nonEmpty?: boolean;
		errors?: string[];
	};

	let { nonEmpty, errors, value = $bindable(''), ...rest }: Props = $props();

	let focused = $state(false);
</script>

<div class={cn('flex flex-col', rest.class)}>
	<input
		bind:value
		onfocus={() => (focused = true)}
		onblur={() => (focused = false)}
		{...rest}
		class="bg-background disabled:text-text/50 w-full rounded-sm border border-zinc-900 px-3 py-2"
	/>
	<FieldError {errors} />
</div>
