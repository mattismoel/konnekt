<script lang="ts">
	import type { Artist } from '$lib/artist';
	import Button from '$lib/components/Button.svelte';
	import Input from '$lib/components/Input.svelte';
	import Selector from '$lib/components/Selector.svelte';
	import type { UpdateConcert } from '$lib/concert';
	import type { Event } from '$lib/event';
	import type { Venue } from '$lib/venue';
	import CreateConcertCard from './CreateConcertCard.svelte';

	type Props = {
		event: Event | null;
		artists: Artist[];
		venues: Venue[];
	};

	const { event, artists, venues }: Props = $props();
	const coverImageUrl = $state(event?.coverImageUrl || '');

	let concerts = $state<(UpdateConcert & { id: string })[]>(
		event?.concerts.map((c) => ({
			id: crypto.randomUUID().toString(),
			artistID: c.artist.id,
			from: c.from,
			to: c.to
		})) || []
	);

	const deleteConcert = (concertId: string) => {
		concerts = concerts.filter((c) => c.id !== concertId);
	};

	const addConcert = () => {
		concerts = [
			...concerts,
			{
				id: crypto.randomUUID(),
				artistID: 0,
				from: new Date(),
				to: new Date()
			}
		];
	};

	$inspect(concerts);
</script>

<form action="" class="space-y-8">
	<h1 class="mb-4 text-2xl font-bold">Lav event.</h1>
	<div class="relative">
		<img src={event?.coverImageUrl || ''} alt="" class="w-full rounded-sm border border-zinc-800" />
		<Button class="absolute right-4 bottom-4">Vælg...</Button>
	</div>
	<div class="flex gap-4 *:flex-1">
		<Input expandX label="Titel" type="text" name="title" />
		<Selector
			class="w-min"
			entries={venues.map((v) => ({
				name: `${v.name}, ${v.city}`,
				value: v.id.toString()
			}))}
		/>
	</div>

	<h1 class="mb-4 text-2xl font-bold">Koncerter.</h1>
	<div class="space-y-4">
		{#each concerts || [] as concert, i (concert.id)}
			<CreateConcertCard
				bind:concert={concerts[i]}
				{artists}
				idx={i + 1}
				onDelete={() => deleteConcert(concert.id)}
			/>
		{/each}
	</div>
	<Button onclick={addConcert} expandX variant="ghost">+ Tilføj koncert</Button>
</form>
