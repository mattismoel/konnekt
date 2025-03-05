<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';
	import Input from './Input.svelte';
	import { format, formatISO } from 'date-fns';

	type Props = HTMLInputAttributes & {
		label: string;
		date: Date;
	};

	let { label, date = $bindable(), ...rest }: Props = $props();

	let value = $derived(format(date, "yyyy-MM-dd'T'HH:mm"));
</script>

<Input
	expandX
	nonEmpty
	type="datetime-local"
	onchange={(e) => (date = e.currentTarget.valueAsDate || new Date())}
	{value}
	{label}
	{...rest}
/>
