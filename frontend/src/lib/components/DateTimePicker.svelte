<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';
	import Input from './Input.svelte';

	type Props = HTMLInputAttributes & {
		label: string;
		date: Date;
	};

	let { label, date = $bindable(), ...rest }: Props = $props();
</script>

<Input
	expandX
	nonEmpty
	type="datetime-local"
	onchange={(e) => (date = e.currentTarget.valueAsDate || new Date())}
	{label}
	{...rest}
/>
