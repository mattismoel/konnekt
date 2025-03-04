<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	type Entry = {
		name: string;
		value: string;
	};

	type Props = HTMLSelectAttributes & {
		selected?: string;
		entries: Entry[];
		onChange: (value: string) => void;
	};

	let { entries, selected, onChange, ...rest }: Props = $props();

	let selectElement: HTMLSelectElement;

	$effect(() => {
		selectElement.value = selected || entries[0].value;
	});
</script>

<select
	onchange={(e) => onChange(e.currentTarget.value)}
	bind:this={selectElement}
	{...rest}
	class={cn('rounded-sm border border-zinc-900 bg-zinc-950', rest.class)}
>
	{#each entries as entry (entry.value)}
		<option value={entry.value}>{entry.name}</option>
	{/each}
</select>
