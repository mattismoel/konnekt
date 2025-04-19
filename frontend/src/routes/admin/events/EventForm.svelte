<script lang="ts">
	import { SvelteMap } from 'svelte/reactivity';

	import type { z } from 'zod';
	import { addMinutes, roundToNearestHours } from 'date-fns';

	import type { Artist } from '$lib/features/artist/artist';
	import type { Venue } from '$lib/features/venue/venue';

	import type { concertForm } from '$lib/features/concert/concert';
	import type { createEventForm, editEventForm, Event } from '$lib/features/event/event';

	import ConcertsList from './ConcertsList.svelte';

	import FieldError from '$lib/components/ui/FieldError.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import ImagePreview from '$lib/components/ImagePreview.svelte';
	import TipTapEditor from '$lib/components/tiptap/TipTapEditor.svelte';
	import Button from '$lib/components/ui/Button.svelte';

	type Props = {
		venues: Venue[];
		artists: Artist[];

		event?: Event;
		errors:
			| z.typeToFlattenedError<z.infer<typeof createEventForm> | z.infer<typeof editEventForm>>
			| undefined;

		onSubmit: (form: z.infer<typeof createEventForm> | z.infer<typeof editEventForm>) => void;
	};

	let { event, venues, artists, errors, onSubmit }: Props = $props();

	const form = $state<z.infer<typeof createEventForm> | z.infer<typeof editEventForm>>(
		event
			? {
					...event,
					venueId: event.venue.id,
					image: null,
					concerts: event.concerts.map(({ artist, from, to }) => ({
						artistID: artist.id,
						from,
						to
					}))
				}
			: {
					title: '',
					description: '',
					concerts: [],
					image: null,
					venueId: 1,
					ticketUrl: ''
				}
	);

	const imageUrl = $derived(form.image ? URL.createObjectURL(form.image) : event?.imageUrl);

	let concerts = $state(
		new SvelteMap<string, z.infer<typeof concertForm>>(
			form.concerts.map((concert) => [crypto.randomUUID(), concert])
		)
	);

	const deleteConcert = (id: string) => {
		concerts.delete(id);
	};

	const addConcert = () => {
		const lastConcert = Array.from(concerts).at(-1)?.[1];

		const from = lastConcert?.to || roundToNearestHours(new Date());
		const to = addMinutes(from, 30);

		const id = crypto.randomUUID();

		concerts.set(id, { from, to, artistID: 1 });
	};

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();

		onSubmit({
			...form,
			concerts: Array.from(concerts.entries()).map(([_, concert]) => concert)
		});
	};
</script>

<form onsubmit={handleSubmit} class="space-y-16">
	<!-- COVER IMAGE -->
	<section>
		<h2 class="mb-8 text-2xl font-semibold">Coverbillede</h2>
		<ImagePreview src={imageUrl} onChange={(file) => (form.image = file)} />
		<FieldError errors={errors?.fieldErrors.image} />
	</section>

	<!-- GENERAL -->
	<section>
		<h2 class="mb-8 text-2xl font-semibold">Generelt</h2>
		<Input
			label="Eventtitel"
			bind:value={form.title}
			class="flex-1"
			errors={errors?.fieldErrors.title}
		/>
		<div class="flex w-full gap-4">
			<Input
				label="Billet-URL"
				class="flex-1"
				bind:value={form.ticketUrl}
				errors={errors?.fieldErrors.ticketUrl}
			/>
			<div>
				<Selector
					class="h-min"
					bind:value={() => form.venueId.toString(), (v) => (form.venueId = parseInt(v))}
					entries={venues.map((venue) => ({ name: venue.name, value: venue.id.toString() }))}
				/>
				<FieldError errors={errors?.fieldErrors.venueId} />
			</div>
		</div>
	</section>

	<!-- EVENT DESCRIPTION -->
	<section>
		<h2 class="mb-8 text-2xl font-semibold">Eventbeskrivelse</h2>
		<div>
			<TipTapEditor bind:value={form.description} />
			<FieldError errors={errors?.fieldErrors.description} />
		</div>
	</section>

	<!-- CONCERTS -->
	<section>
		<h2 class="mb-8 text-2xl font-semibold">Koncerter</h2>
		<ConcertsList {concerts} {artists} onAdd={addConcert} onDelete={deleteConcert} />
		<FieldError errors={errors?.fieldErrors.concerts} />
	</section>

	<div class="flex flex-col gap-2 md:flex-row">
		<Button variant="ghost" class="w-full md:max-w-64">Preview</Button>
		<Button type="submit" class="w-full md:max-w-64">Offentliggør</Button>
	</div>
</form>
