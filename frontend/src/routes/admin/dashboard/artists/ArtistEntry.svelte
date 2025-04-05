<script lang="ts">
	import type { Artist } from '$lib/artist';
	import { hasPermissions, type Permission } from '$lib/auth';
	import Button from '$lib/components/ui/Button.svelte';
	import TrashIcon from '~icons/mdi/trash';

	type Props = {
		artist: Artist;
		onDelete: () => void;
		userPermissions: Permission[];
	};

	let { artist, userPermissions, onDelete }: Props = $props();
</script>

<li>
	<a
		href="/admin/artists/edit?id={artist.id}"
		class="flex items-center rounded-sm border border-transparent px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
	>
		<span class="flex-1 font-medium">{artist.name}</span>
		<Button
			disabled={!hasPermissions(userPermissions, ['delete:artist'])}
			variant="dangerous"
			onclick={onDelete}><TrashIcon /></Button
		>
	</a>
</li>
