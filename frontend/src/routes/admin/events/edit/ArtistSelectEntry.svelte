<script lang="ts">
	import type { Artist } from '$lib/features/artist/artist';

	import SelectEntry from '$lib/components/SelectEntry.svelte';

	import OpenIcon from '~icons/mdi/open-in-new';

	type Props = {
		artist: Artist;
		onSelect: () => void;
		onDeselect: () => void;
		selected: boolean;
	};

	let { artist, selected, onSelect, onDeselect }: Props = $props();
</script>

<SelectEntry {selected} {onSelect} {onDeselect}>
	<span class="flex-1 text-left">{artist.name}</span>
	<a class="hover:text-text text-zinc-500" target="_blank" href="/admin/artists/edit?id={artist.id}"
		><OpenIcon /></a
	>
</SelectEntry>
