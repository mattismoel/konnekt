<script lang="ts">
	import { DATE_FORMAT } from '$lib/time';

	import type { Event } from '$lib/event';

	import { earliestConcert, latestConcert } from '$lib/concert';
	import { format, isBefore, startOfToday, startOfTomorrow, startOfYesterday } from 'date-fns';

	import MenuIcon from '~icons/mdi/dots-vertical';
	import Button from '$lib/components/ui/Button.svelte';
	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import { goto } from '$app/navigation';

	type Props = {
		event: Event;
		onDelete: () => void;
	};

	let { event, onDelete }: Props = $props();

	const fromDate = $derived(earliestConcert(event.concerts)?.from || new Date());
	const toDate = $derived(latestConcert(event.concerts)?.to || new Date());

	let showContextMenu = $state(false);
	let contextBtn: HTMLButtonElement;
</script>

<li
	class:expired={isBefore(fromDate, startOfYesterday())}
	class="group relative flex rounded-md border border-zinc-900 px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
>
	<a class="flex flex-1 items-center gap-4" href="/admin/events/edit?id={event.id}">
		<span
			class="group-[.expired]:text-text/50 flex-1 font-medium group-[.expired]:italic group-[.expired]:line-through"
			>{event.title}</span
		>
		<span class="text-text/50 line-clamp-1 hidden flex-1 md:inline"
			>{format(fromDate, DATE_FORMAT)}</span
		>
		<span class="text-text/50 line-clamp-1 hidden flex-1 md:inline"
			>{format(fromDate, 'HH:mm')} - {format(toDate, 'HH:mm')}</span
		>
	</a>
	<button
		bind:this={contextBtn}
		onclick={() => (showContextMenu = !showContextMenu)}
		class="text-text/50 hover:text-text"
	>
		<MenuIcon />
	</button>
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry action={() => goto(`/admin/events/edit?id=${event.id}`)}
			>Redigér</ContextMenuEntry
		>
		<ContextMenuEntry action={onDelete}>Slet</ContextMenuEntry>
	</ContextMenu>
</li>
