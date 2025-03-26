<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';
	import TableCell from './TableCell.svelte';

	type Props = HTMLAttributes<HTMLInputElement> & {
		value?: string;
	};

	let { value = $bindable(), ...rest }: Props = $props();
</script>

<TableCell>
	<input
		bind:value
		type="text"
		class="w-full rounded-md border border-zinc-900 bg-zinc-950"
		{...rest}
	/>
</TableCell>
