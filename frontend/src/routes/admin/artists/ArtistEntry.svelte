<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { deleteArtist, type Artist } from '$lib/features/artist/artist';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import MenuIcon from '~icons/mdi/dots-vertical';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { APIError } from '$lib/api';

	type Props = {
		artist: Artist;
		memberPermissions: Permission[];
	};

	let { artist, memberPermissions }: Props = $props();

	let showContextMenu = $state(false);

	const handleDelete = async () => {
		if (!confirm(`Er du sikke på, at du vil slette ${artist.name}?`)) return;

		try {
			await deleteArtist(fetch, artist.id);
			toaster.addToast('Kunstner slettet');
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke slette kunstner', e.cause, 'error');
				throw e;
			}

			toaster.addToast('Kunne ikke slette kunstner', 'Noget gik galt...', 'error');
			throw e;
		}
	};
</script>

<ListEntry>
	<a href="/admin/artists/edit/{artist.id}" class="flex w-full flex-col">
		<span>{artist.name}</span>
		<span class="text-text/50">{artist.genres.map((genre) => genre.name).join(', ')}</span>
	</a>
	<button onclick={() => (showContextMenu = true)} class="rounded-md p-2 hover:bg-zinc-900">
		<MenuIcon />
	</button>
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['edit:artist'])}
			href="/admin/artists/edit/{artist.id}"
		>
			Redigér
		</ContextMenuEntry>
		<ContextMenuEntry
			onclick={handleDelete}
			disabled={!hasPermissions(memberPermissions, ['delete:artist'])}
		>
			Slet
		</ContextMenuEntry>
	</ContextMenu>
</ListEntry>
