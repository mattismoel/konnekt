<script lang="ts">
	import type { Artist } from '$lib/artist';
	import Button from '$lib/components/Button.svelte';
	import Input from '$lib/components/Input.svelte';
	import Selector from '$lib/components/Selector.svelte';
	import type { Event, eventForm } from '$lib/event';
	import type { Venue } from '$lib/venue';
	import { z } from 'zod';
	import CreateConcertCard from './CreateConcertCard.svelte';

	type Props = {
		event: Event | null;
		artists: Artist[];
		venues: Venue[];

		onSubmit: (e: z.infer<typeof eventForm>) => void;
	};

	const { event: prevEvent, artists, venues, onSubmit }: Props = $props();

	let event = $state<z.infer<typeof eventForm>>({
		title: prevEvent?.title || '',
		description: prevEvent?.description || '',
		venueId: prevEvent?.venue.id || 1,
		coverImageUrl: prevEvent?.coverImageUrl || '',
		concerts:
			prevEvent?.concerts.map((c) => ({
				idxId: crypto.randomUUID(),
				artistID: c.artist.id,
				from: c.from,
				to: c.to
			})) || []
	});

	const deleteConcert = (idxId: string) => {
		event.concerts = event.concerts.filter((c, i) => c.idxId !== idxId);
	};

	const addConcert = () => {
		event.concerts = [
			...event.concerts,
			{
				idxId: crypto.randomUUID(),
				artistID: 1,
				from: new Date(),
				to: new Date()
			}
		];
	};

	const submit = (e: SubmitEvent & { currentTarget: EventTarget & HTMLFormElement }) => {
		e.preventDefault();
		onSubmit(event);
	};

	$inspect(event.concerts);
</script>

<form class="space-y-8" onsubmit={submit}>
	<h1 class="mb-4 text-2xl font-bold">Lav event.</h1>
	<div class="relative">
		<img src={event.coverImageUrl} alt="" class="w-full rounded-sm border border-zinc-800" />
		<Button class="absolute right-4 bottom-4">Vælg...</Button>
	</div>
	<div class="flex gap-4 *:flex-1">
		<Input expandX label="Titel" type="text" name="title" value={event.title} />
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
		{#each event.concerts || [] as concert, i (concert.idxId)}
			<CreateConcertCard
				bind:concert={event.concerts[i]}
				{artists}
				idx={i + 1}
				onDelete={() => deleteConcert(concert.idxId)}
			/>
		{/each}
	</div>
	<Button onclick={addConcert} expandX variant="ghost">+ Tilføj koncert</Button>
	<div class="flex gap-4">
		<Button expandX variant="secondary">Preview</Button>
		<Button type="submit" expandX>Offentligør</Button>
	</div>
</form>
