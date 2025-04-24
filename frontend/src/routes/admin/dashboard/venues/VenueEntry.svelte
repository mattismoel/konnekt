<script lang="ts">
	import { cn } from '$lib/clsx';
	import Selector from '$lib/components/ui/Selector.svelte';
	import { COUNTRIES_MAP } from '$lib/features/venue/countries';
	import { deleteVenue, editVenue, type Venue, type venueForm } from '$lib/features/venue/venue';
	import type { z } from 'zod';

	import CheckIcon from '~icons/mdi/check';
	import MenuIcon from '~icons/mdi/dots-vertical';
	import XIcon from '~icons/mdi/close';
	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';
	import { APIError } from '$lib/api';

	type Props = {
		venue: Venue;
		memberPermissions: Permission[];
	};

	let { venue, memberPermissions }: Props = $props();

	let showContextMenu = $state(false);

	let form = $state<z.infer<typeof venueForm>>({
		...(venue || {
			name: '',
			city: '',
			countryCode: ''
		})
	});

	let isEdited = $derived(
		form.name !== venue?.name ||
			form.city !== venue?.city ||
			form.countryCode !== venue?.countryCode
	);

	const resetForm = () => {
		form = { ...(venue || { name: '', city: '', countryCode: '' }) };
	};

	const handleEditVenue = async () => {
		try {
			await editVenue(fetch, venue.id, form);
			toaster.addToast('Venue redigeret');
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke regigere venue', e.cause, 'error');
				return;
			}

			toaster.addToast('Kunne ikke redigere venue', 'Noget gik galt...', 'error');
			return;
		}
	};

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

<svelte:window
	onkeydown={(e) => {
		if (e.key === 'Escape' && isEdited) resetForm();
	}}
/>

<ListEntry
	class={cn(
		'group relative flex items-center gap-8 rounded-md border border-transparent p-2 hover:border-zinc-800 hover:bg-zinc-900',
		{
			'border-green-800 bg-green-950 hover:border-green-800 hover:bg-green-950': isEdited
		}
	)}
>
	<input
		bind:value={form.name}
		class={cn('w-full rounded-sm border-transparent bg-transparent font-semibold', {
			italic: form.name !== venue?.name
		})}
	/>
	<input
		bind:value={form.city}
		class={cn('text-text/50 hidden w-full rounded-sm border-transparent bg-transparent lg:inline', {
			italic: form.city !== venue?.city
		})}
	/>
	<Selector
		class={cn(
			'text-text/50 hidden border-transparent bg-transparent group-hover:border-zinc-800 lg:inline',
			{
				italic: form.countryCode !== venue?.countryCode
			}
		)}
		entries={Array.from(COUNTRIES_MAP).map(([key, val]) => ({ name: val, value: key }))}
		bind:value={form.countryCode}
	/>
	<div class="text-text/50 flex w-full justify-end gap-8">
		{#if isEdited}
			<div class="flex gap-4">
				<button onclick={resetForm} class="hover:text-red-400">
					<XIcon />
				</button>
				<button onclick={handleEditVenue} class="hover:text-green-400"><CheckIcon /></button>
			</div>
		{/if}
	</div>
	<button onclick={() => (showContextMenu = true)}>
		<MenuIcon />
	</button>
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:venue'])}
			action={handleDeleteVenue}>Slet</ContextMenuEntry
		>
	</ContextMenu>
</ListEntry>
