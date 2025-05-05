<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';
	import Input from './ui/Input.svelte';
	import { formatISO } from 'date-fns';

	type Props = Omit<HTMLInputAttributes, 'defaultValue'> & {
		defaultValue: Date;
		onChange: (newDate: Date) => void;
	};

	let { defaultValue, onChange, ...rest }: Props = $props();

	let dateString = $state(formatISO(defaultValue).slice(0, 16));

	const handleChange = (e: Event) => {
		const currentTarget = e.currentTarget as HTMLInputElement;
		const newDate = new Date(currentTarget.value);
		onChange(newDate);
	};
</script>

<Input type="datetime-local" value={dateString} onchange={handleChange} {...rest} />
