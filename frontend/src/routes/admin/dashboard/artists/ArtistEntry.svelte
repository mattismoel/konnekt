<script lang="ts">
	import type { Artist } from '$lib/artist';
	import Button from '$lib/components/ui/Button.svelte';
	import TrashIcon from '~icons/mdi/trash';

	type Props = {
		artist: Artist;
		onDelete: () => void;
	};

	let { artist, onDelete }: Props = $props();
</script>

<li>
	<a
		href="/admin/artists/edit?id={artist.id}"
		class="flex items-center rounded-sm border border-transparent px-4 py-2 hover:border-zinc-700 hover:bg-zinc-800"
	>
		<span class="flex-1 font-medium">{artist.name}</span>
		<Button variant="dangerous" onclick={onDelete}><TrashIcon /></Button>
	</a>
</li>
