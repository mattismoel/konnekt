<script lang="ts">
	import PlusIcon from '~icons/mdi/plus';
	import Button from '$lib/components/ui/Button.svelte';

	import CleanIcon from '~icons/mdi/broom';

	import { goto, invalidateAll } from '$app/navigation';
	import SearchBar from '$lib/components/SearchBar.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';
	import EventList from './EventList.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { deleteEvent } from '$lib/features/event/event';
	import { tryCatch } from '$lib/error';
	import { APIError } from '$lib/api';
	import DashboardLayout from '../DashboardLayout.svelte';
	import DashboardHeader from '../DashboardHeader.svelte';
	import HeaderActions from '../HeaderActions.svelte';

	let { data } = $props();

	let search = $state('');

	let filteredEvents = $derived(
		[...data.previousEvents, ...data.upcomingEvents].filter((event) =>
			event.title.toLowerCase().includes(search.toLowerCase())
		)
	);

	const handleCleanPreviousEvents = async () => {
		let message = 'Er du sikker på, at du vil rydde alle tidligere events?\n\n';

		data.previousEvents.map((e) => {
			message += `- ${e.title}\n`;
		});

		message += '\n';
		message += 'Handlingen kan ikke fortrydes.';

		if (!confirm(message)) return;

		data.previousEvents.forEach(async ({ id }) => {
			const { error } = await tryCatch(deleteEvent(fetch, id));
			if (error) {
				if (error instanceof APIError) {
					toaster.addToast('Kunne ikke rydde op', error.cause, 'error');
					return;
				}

				toaster.addToast('Kunne ikke rydde op', 'Noget gik galt...', 'error');
				return;
			}
		});

		toaster.addToast('Events blev ryddet op');
		await invalidateAll();
	};
</script>

<DashboardLayout>
	<DashboardHeader title="Events" description="Overblik over alle events.">
		<HeaderActions>
			<Button
				disabled={!hasPermissions(data.member.permissions, ['edit:event'])}
				onclick={() => goto(`/admin/events/create`)}
			>
				<PlusIcon />Tilføj
			</Button>
			<Button onclick={handleCleanPreviousEvents} variant="ghost"><CleanIcon />Ryd</Button>
		</HeaderActions>
	</DashboardHeader>

	{#if hasPermissions(data.member.permissions, ['view:event'])}
		<div class="space-y-8">
			<SearchBar bind:value={search} />
			{#if search.trim() !== ''}
				<section>
					<EventList events={filteredEvents} memberPermissions={data.member.permissions} />
				</section>
			{:else}
				<section>
					<EventList events={data.upcomingEvents} memberPermissions={data.member.permissions} />
				</section>
			{/if}

			{#if data.previousEvents.length > 0}
				<section>
					<details>
						<summary class="mb-4">Tidligere events ({data.previousEvents.length})</summary>
						<EventList events={data.previousEvents} memberPermissions={data.member.permissions} />
					</details>
				</section>
			{/if}
		</div>
	{:else}
		<span>Du har ikke tilladelse til at se denne side...</span>
	{/if}
</DashboardLayout>
