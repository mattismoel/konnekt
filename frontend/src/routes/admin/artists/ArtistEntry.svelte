<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import { deleteArtist, type Artist } from '$lib/features/artist/artist';
	import { hasPermissions } from '$lib/features/auth/permission';
	import * as ContextMenu from '$lib/components/ui/context-menu';
	import * as List from '$lib/components/ui/list/index';
	import { toaster } from '$lib/toaster.svelte';
	import { APIError } from '$lib/api';
	import { authStore } from '$lib/auth.svelte';

	type Props = {
		artist: Artist;
	};

	let { artist }: Props = $props();

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

<List.Entry>
	<List.Section href="/admin/artists/edit/{artist.id}">
		<span>{artist.name}</span>
		<span class="text-text/50">{artist.genres.map((genre) => genre.name).join(', ')}</span>
	</List.Section>
	<List.Section expand={false}>
		<ContextMenu.Button onclick={() => (showContextMenu = true)} />
	</List.Section>
	<ContextMenu.Root bind:show={showContextMenu} class="absolute top-1/2 right-4">
		<ContextMenu.Entry
			disabled={!hasPermissions(authStore.permissions, ['edit:artist'])}
			href="/admin/artists/edit/{artist.id}"
		>
			Redigér
		</ContextMenu.Entry>
		<ContextMenu.Entry
			onclick={handleDelete}
			disabled={!hasPermissions(authStore.permissions, ['delete:artist'])}
		>
			Slet
		</ContextMenu.Entry>
	</ContextMenu.Root>
</List.Entry>
