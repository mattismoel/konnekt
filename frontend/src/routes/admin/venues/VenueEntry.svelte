<script lang="ts">
	import { deleteVenue, type Venue } from '$lib/features/venue/venue';

	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';
	import { APIError } from '$lib/api';
	import ContextMenuButton from '$lib/components/ui/context-menu/ContextMenuButton.svelte';

	type Props = {
		venue: Venue;
		memberPermissions: Permission[];
	};

	let { venue, memberPermissions }: Props = $props();
	let showContextMenu = $state(false);

	const handleDeleteVenue = async () => {
		if (!confirm(`Er sikker på, at du vil slette venue "${venue.name}"?`)) return;

		try {
			await deleteVenue(fetch, venue.id);
			toaster.addToast('Venue slettet');
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke slette venue', e.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke slette venue', 'Noget gik galt...', 'error');
			return;
		}
	};
</script>

<ListEntry class="">
	<a href="/admin/venues/edit/{venue.id}" class="w-full">
		<div class="flex flex-col">
			<span>{venue.name}</span>
			<span class="text-text/50">{venue.city}, {venue.countryCode}</span>
		</div>
	</a>

	<ContextMenuButton onclick={() => (showContextMenu = true)} />
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:venue'])}
			href="/admin/venues/edit/{venue.id}">Redigér</ContextMenuEntry
		>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:venue'])}
			onclick={handleDeleteVenue}>Slet</ContextMenuEntry
		>
	</ContextMenu>
</ListEntry>
