<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLSelectAttributes } from 'svelte/elements';

	type Entry = {
		name: string;
		value: string;
	};

	type Props = HTMLSelectAttributes & {
		entries: Entry[];
	};

	let { entries, children, ...rest }: Props = $props();
</script>

<select {...rest} class={cn('rounded-sm border border-zinc-900 bg-zinc-950', rest.class)}>
	{#each entries as entry (entry.value)}
		<option value={entry.value}>{entry.name}</option>
	{/each}
</select>
