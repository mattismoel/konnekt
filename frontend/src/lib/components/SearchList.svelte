<script lang="ts">
	import type { WithChildren } from '$lib/types';
	import SearchBar from './SearchBar.svelte';
	import List from './ui/List.svelte';

	type Props = WithChildren<{
		search: string;
	}>;

	let { search = $bindable(''), children }: Props = $props();
</script>

<div class="flex flex-col gap-8">
	<SearchBar bind:value={search} />
	<List>
		{@render children?.()}
	</List>
</div>
