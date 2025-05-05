<script lang="ts">
	import { deleteVenue, type Venue } from '$lib/features/venue/venue';

	import * as ContextMenu from '$lib/components/ui/context-menu/index';
	import { hasPermissions } from '$lib/features/auth/permission';
	import * as List from '$lib/components/ui/list/index';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';
	import { APIError } from '$lib/api';
	import { authStore } from '$lib/auth.svelte';

	type Props = {
		venue: Venue;
	};

	let { venue }: Props = $props();
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

<List.Entry>
	<List.Section href="/admin/venues/edit/{venue.id}">
		<span>{venue.name}</span>
		<span class="text-text/50">{venue.city}, {venue.countryCode}</span>
	</List.Section>

	<List.Section expand={false}>
		<ContextMenu.Button onclick={() => (showContextMenu = true)} />
	</List.Section>

	<ContextMenu.Root bind:show={showContextMenu} class="absolute top-1/2 right-4">
		<ContextMenu.Entry
			disabled={!hasPermissions(authStore.permissions, ['delete:venue'])}
			href="/admin/venues/edit/{venue.id}"
		>
			Redigér
		</ContextMenu.Entry>
		<ContextMenu.Entry
			disabled={!hasPermissions(authStore.permissions, ['delete:venue'])}
			onclick={handleDeleteVenue}
		>
			Slet
		</ContextMenu.Entry>
	</ContextMenu.Root>
</List.Entry>
