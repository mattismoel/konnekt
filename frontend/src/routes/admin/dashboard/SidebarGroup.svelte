<script lang="ts">
	import Accordion from '$lib/components/ui/Accordion.svelte';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		label: string;
	};

	let { label, children }: Props = $props();
</script>

<div>
	<h3 class="font-heading mb-4 font-semibold">{label}</h3>
</div>
