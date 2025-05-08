<script lang="ts">
	import { DATETIME_FORMAT } from '$lib/time';

	import { deleteEvent, type Event } from '$lib/features/event/event';

	import { earliestConcert } from '$lib/features/concert/concert';
	import { format, isBefore, startOfToday } from 'date-fns';

	import * as ContextMenu from '$lib/components/ui/context-menu/index';
	import { invalidateAll } from '$app/navigation';
	import { hasPermissions } from '$lib/features/auth/permission';
	import * as List from '$lib/components/ui/list/index';
	import { toaster } from '$lib/toaster.svelte';
	import { APIError } from '$lib/api';

	import LocationIcon from '~icons/mdi/location';
	import { authStore } from '$lib/auth.svelte';

	type Props = {
		event: Event;
	};

	let { event }: Props = $props();

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

<List.Entry title="Redigér event" class={`group ${expired ? 'expired' : ''}`}>
	<List.Section href="/admin/events/edit/{event.id}">
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
	</List.Section>
	<List.Section href="/admin/venues/edit/{event.venue.id}" class="group/venue">
		<span
			class:disabled={!hasPermissions(authStore.permissions, ['edit:venue'])}
			class="text-text/50 group-hover/venue:text-text group-[.disabled]/venue:text-text/50 hidden w-full items-center gap-2 group-hover/venue:underline group-[.disabled]/venue:no-underline md:flex"
		>
			<LocationIcon />
			<span class="whitespace-nowrap">{event.venue.name}</span>
		</span>
	</List.Section>

	<List.Section expand={false}>
		<ContextMenu.Button onclick={() => (showContextMenu = true)} />
	</List.Section>

	<ContextMenu.Root bind:show={showContextMenu} class="absolute top-1/2 right-4">
		<ContextMenu.Entry
			disabled={!hasPermissions(authStore.permissions, ['edit:event'])}
			href="/admin/events/edit/{event.id}"
		>
			Redigér
		</ContextMenu.Entry>
		<ContextMenu.Entry
			disabled={!hasPermissions(authStore.permissions, ['delete:event'])}
			onclick={handleDeleteEvent}
		>
			Slet
		</ContextMenu.Entry>
	</ContextMenu.Root>
</List.Entry>
