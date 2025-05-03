<script lang="ts">
	import { DATETIME_FORMAT } from '$lib/time';

	import { deleteEvent, type Event } from '$lib/features/event/event';

	import { earliestConcert } from '$lib/features/concert/concert';
	import { format, isBefore, startOfToday } from 'date-fns';

	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import { invalidateAll } from '$app/navigation';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import ContextMenuButton from '$lib/components/ui/context-menu/ContextMenuButton.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { APIError } from '$lib/api';

	import LocationIcon from '~icons/mdi/location';

	type Props = {
		event: Event;
		memberPermissions: Permission[];
	};

	let { event, memberPermissions }: Props = $props();

	let showContextMenu = $state(false);

	let artists = $derived(event.concerts.map((concert) => concert.artist));

	const fromDate = $derived(earliestConcert(event.concerts)?.from || new Date());
	let expired = $derived(isBefore(fromDate, startOfToday()));

	const handleDeleteEvent = async () => {
		if (!confirm(`Vil du slette ${event.title}?`)) return;

		try {
			await deleteEvent(fetch, event.id);
			toaster.addToast('Event slettet');
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke slette event', e.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke slette event', 'Noget gik galt...', 'error');
			return;
		}
	};
</script>

<ListEntry class={`group ${expired ? 'expired' : ''}`}>
	<a title="Redigér event" href="/admin/events/edit/{event.id}" class="flex w-full flex-col">
		<span class="line-clamp-1 group-[.expired]:line-through">{event.title}</span>
		<span class="text-text/50 line-clamp-1">{format(fromDate, DATETIME_FORMAT)}</span>
		<span class="text-text/50 line-clamp-1 md:hidden">{event.venue.name}</span>
		<span class="text-text/50 line-clamp-1 hidden md:block">
			{#if artists.length > 2}
				{artists
					.slice(0, 2)
					.map((artist) => artist.name)
					.join(', ')} (+{artists.length - 2} mere)
			{:else}
				{artists.map((artist) => artist.name).join(', ')}
			{/if}
		</span>
	</a>
	<a
		title="Redigér venue"
		href="/admin/venues/edit/{event.venue.id}"
		class="text-text/50 hover:text-text hidden w-full items-center gap-2 hover:underline md:flex"
	>
		<LocationIcon />
		<span class="whitespace-nowrap">{event.venue.name}</span>
	</a>
	<ContextMenuButton onclick={() => (showContextMenu = true)} />
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['edit:event'])}
			href="/admin/events/edit/{event.id}"
		>
			Redigér
		</ContextMenuEntry>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:event'])}
			onclick={handleDeleteEvent}>Slet</ContextMenuEntry
		>
	</ContextMenu>
</ListEntry>
