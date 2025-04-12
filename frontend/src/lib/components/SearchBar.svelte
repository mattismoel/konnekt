<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';
	import SearchIcon from '~icons/mdi/search';

	type Props = HTMLAttributes<HTMLInputElement> & {
		value: string;
	};

	let { value = $bindable(''), placeholder = 'Søg...', ...rest }: Props = $props();
</script>

<div class="relative w-full">
	<input
		{...rest}
		{placeholder}
		bind:value
		type="text"
		class="h-full w-full min-w-48 rounded-sm border-zinc-800 bg-zinc-900 py-2 pr-4 pl-12"
	/>
	<SearchIcon class="text-text/75 absolute top-1/2 left-4 -translate-y-1/2" />
</div>
