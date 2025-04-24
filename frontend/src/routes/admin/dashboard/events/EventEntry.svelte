<script lang="ts">
	import { DATE_FORMAT } from '$lib/time';

	import type { Event } from '$lib/features/event/event';

	import { earliestConcert } from '$lib/features/concert/concert';
	import { format, isBefore, startOfToday } from 'date-fns';

	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import { goto } from '$app/navigation';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import ContextMenuButton from '$lib/components/ui/context-menu/ContextMenuButton.svelte';

	type Props = {
		event: Event;
		memberPermissions: Permission[];
		onDelete: () => void;
	};

	let { event, memberPermissions, onDelete }: Props = $props();

	let showContextMenu = $state(false);

	let artists = $derived(event.concerts.map((concert) => concert.artist));

	const fromDate = $derived(earliestConcert(event.concerts)?.from || new Date());
	let expired = $derived(isBefore(fromDate, startOfToday()));
</script>

<ListEntry class={`group ${expired ? 'expired' : ''}`}>
	<a href="/admin/events/edit/{event.id}">
		<div class="flex flex-1 flex-col">
			<span class="line-clamp-1 group-[.expired]:line-through">{event.title}</span>
			<span class="text-text/50">{format(fromDate, DATE_FORMAT)}</span>
		</div>
		<span class="text-text/50 hidden sm:block">
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
	<ContextMenuButton onclick={() => (showContextMenu = true)} />
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['edit:event'])}
			action={() => goto(`/admin/events/edit?id=${event.id}`)}>Redigér</ContextMenuEntry
		>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:event'])}
			action={onDelete}>Slet</ContextMenuEntry
		>
	</ContextMenu>
</ListEntry>
