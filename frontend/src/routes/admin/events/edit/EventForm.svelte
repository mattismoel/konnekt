<script lang="ts">
	import type { Artist } from '$lib/artist';
	import Button from '$lib/components/ui/Button.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import { eventForm, type Event } from '$lib/event';
	import type { Venue } from '$lib/venue';
	import { z, ZodError } from 'zod';
	import CreateConcertCard from './CreateConcertCard.svelte';
	import ImagePreview from '$lib/components/ui/ImagePreview.svelte';
	import { concertForm } from '$lib/concert';

	type Props = {
		event: Event | null;
		artists: Artist[];
		venues: Venue[];

		onSubmit: (e: z.infer<typeof eventForm>) => void;
	};

	const { event, artists, venues, onSubmit }: Props = $props();

	const concertWithID = concertForm.extend({ id: z.string().uuid() });

	const extendedForm = eventForm.extend({
		concerts: concertWithID.array()
	});

	let form = $state<z.infer<typeof extendedForm>>({
		title: event?.title || '',
		description: event?.description || '',
		venueId: event?.venue.id || 1,
		coverImage: null,
		concerts:
			event?.concerts.map((c) => ({
				id: crypto.randomUUID().toString(),
				artistID: c.artist.id,
				from: c.from,
				to: c.to
			})) || []
	});

	let formError = $state<ZodError>();
	let coverImageUrl = $derived(
		form.coverImage ? URL.createObjectURL(form.coverImage) : event?.coverImageUrl || ''
	);

	const deleteConcert = (id: string) => {
		form.concerts = form.concerts.filter((c) => c.id !== id);
	};

	const addConcert = () => {
		form.concerts = [
			...form.concerts,
			{ id: crypto.randomUUID().toString(), artistID: 1, from, to }
		];
	};

	const submit = (e: SubmitEvent & { currentTarget: EventTarget & HTMLFormElement }) => {
		e.preventDefault();

		const { error, data, success } = eventForm.safeParse(form);
		if (!success) {
			formError = error;
			return;
		}

		onSubmit(data);
	};
</script>

<form class="space-y-8" onsubmit={submit}>
	<h1 class="mb-4 text-2xl font-bold">Lav event.</h1>
	<ImagePreview src={coverImageUrl || ''} onChange={(file) => (form.coverImage = file)} />
	<div class="flex gap-4 *:flex-1">
		<Input expandX label="Titel" type="text" name="title" value={form.title} />
		<Selector
			class="w-min"
			selected={form.venueId}
			entries={venues.map((v) => ({
				name: `${v.name}, ${v.city}`,
				value: v.id.toString()
			}))}
		/>
	</div>

	<h1 class="mb-4 text-2xl font-bold">Koncerter.</h1>
	<div class="space-y-4">
		{#each form.concerts || [] as concert, i (concert.id)}
			<CreateConcertCard
				bind:concert={form.concerts[i]}
				{artists}
				idx={i + 1}
				onDelete={() => deleteConcert(concert.id)}
			/>
		{/each}
	</div>
	<Button onclick={addConcert} expandX variant="ghost">+ Tilføj koncert</Button>
	<div class="flex gap-4">
		<Button expandX variant="secondary">Preview</Button>
		<Button type="submit" expandX>Offentligør</Button>
	</div>
</form>
