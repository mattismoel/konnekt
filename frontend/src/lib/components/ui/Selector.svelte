<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	type Entry = {
		name: string;
		value: string;
	};

	type Props = Omit<HTMLSelectAttributes, 'value'> & {
		entries: Entry[];
		value?: string | undefined | null;
	};

	let { entries, value = $bindable(), placeholder = 'Vælg...', ...rest }: Props = $props();

	let selectElement: HTMLSelectElement;

	$effect(() => {
		selectElement.value = value || entries[0].value;
	});
</script>

<select
	bind:this={selectElement}
	bind:value
	{...rest}
	class={cn('rounded-sm border border-zinc-900 bg-zinc-950', rest.class)}
>
	<option value="" disabled selected={!value}>{placeholder}</option>
	{#each entries as entry (entry.value)}
		<option value={entry.value}>{entry.name}</option>
	{/each}
</select>
