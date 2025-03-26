<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';
	import Input from './Input.svelte';
	import { formatHTMLDatetimeLocal } from '$lib/time';
	import { formatISO } from 'date-fns';

	type Props = Omit<HTMLInputAttributes, 'defaultValue'> & {
		label: string;
		defaultValue: Date;
		onChange: (newDate: Date) => void;
	};

	let { label, defaultValue, onChange, ...rest }: Props = $props();

	let dateString = $state(formatISO(defaultValue).slice(0, 16));

	const handleChange = (e: Event) => {
		const currentTarget = e.currentTarget as HTMLInputElement;
		const newDate = new Date(currentTarget.value);
		onChange(newDate);
	};
</script>

<Input
	nonEmpty
	type="datetime-local"
	value={dateString}
	onchange={handleChange}
	{label}
	{...rest}
/>
