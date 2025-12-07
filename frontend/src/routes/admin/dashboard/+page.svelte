<script lang="ts">
	import { getUpcomingEvents } from "$lib/features/events/event.remote";
	import EventListEntry from "$lib/features/events/components/EventListEntry.svelte";
	import ArtistListEntry from "$lib/features/artists/components/ArtistListEntry.svelte";

	const { items: upcomingEvents } = await getUpcomingEvents(undefined);

	const upcomingArtists = $derived(
		upcomingEvents.flatMap(({ concerts }) => concerts.flatMap((concert) => concert.artist))
	);
</script>

<main class="mx-responsive py-32">
	<h1 class="font-heading mb-8 text-4xl font-bold">Dashboard</h1>

	<div class="grid grid-cols-2 gap-16">
		<div>
			<h1 class="font-heading mb-4 text-2xl font-bold">Kommende events</h1>
			<ul>
				{#each upcomingEvents as event}
					<EventListEntry {event} />
				{/each}
			</ul>
		</div>

		<div>
			<h1 class="font-heading mb-4 text-2xl font-bold">Kommende kunstnere</h1>
			<ul class="flex flex-col gap-2">
				{#each upcomingArtists as artist}
					<ArtistListEntry {artist} />
				{/each}
			</ul>
		</div>
	</div>
</main>
