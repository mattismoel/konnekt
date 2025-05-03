<script lang="ts">
	import type { HTMLInputAttributes } from 'svelte/elements';

	type Props = HTMLInputAttributes;

	let { value = $bindable(), ...rest }: Props = $props();
</script>

<input
	{...rest}
	bind:value
	class="bg-background disabled:text-text/50 w-full rounded-sm border border-zinc-900 px-3 py-2"
/>
