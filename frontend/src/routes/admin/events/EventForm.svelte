<script lang="ts">
	import { SvelteMap } from 'svelte/reactivity';

	import type { z } from 'zod';
	import { addMinutes, roundToNearestHours } from 'date-fns';

	import type { Artist } from '$lib/features/artist/artist';
	import type { Venue } from '$lib/features/venue/venue';

	import type { concertForm } from '$lib/features/concert/concert';
	import type { createEventForm, editEventForm, Event } from '$lib/features/event/event';

	import ConcertsList from './ConcertsList.svelte';
	import PublishIcon from '~icons/mdi/upload';
	import AddIcon from '~icons/mdi/add';

	import Input from '$lib/components/ui/Input.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import ImagePreview from '$lib/components/ImagePreview.svelte';
	import TipTapEditor from '$lib/components/tiptap/TipTapEditor.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Spinner from '$lib/components/Spinner.svelte';
	import FormField from '$lib/components/ui/FormField.svelte';

	type Props = {
		venues: Venue[];
		artists: Artist[];

		event?: Event;
		errors:
			| z.typeToFlattenedError<z.infer<typeof createEventForm> | z.infer<typeof editEventForm>>
			| undefined;

		loading: boolean;

		onSubmit: (form: z.infer<typeof createEventForm> | z.infer<typeof editEventForm>) => void;
	};

	let { event, venues, artists, errors, loading, onSubmit }: Props = $props();

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
		<FormField errors={errors?.fieldErrors.image}>
			<ImagePreview
				accept="image/jpeg,image/png"
				src={imageUrl}
				onChange={(file) => (form.image = file)}
			/>
		</FormField>
	</section>

	<!-- GENERAL -->
	<section>
		<h2 class="mb-8 text-2xl font-semibold">Generelt</h2>

		<div class="space-y-4">
			<FormField errors={errors?.fieldErrors.title}>
				<Input placeholder="Eventtitel" bind:value={form.title} class="flex-1" />
			</FormField>

			<div class="flex w-full gap-4">
				<FormField errors={errors?.fieldErrors.ticketUrl}>
					<Input placeholder="Billet-URL" class="flex-1" bind:value={form.ticketUrl} />
				</FormField>

				<FormField errors={errors?.fieldErrors.venueId}>
					<div class="flex gap-4">
						<Selector
							class="w-full"
							bind:value={() => form.venueId.toString(), (v) => (form.venueId = parseInt(v))}
							entries={venues.map((venue) => ({ name: venue.name, value: venue.id.toString() }))}
						/>
						<Button target="__blank" variant="secondary" href="/admin/venues/create"
							><AddIcon />Ny</Button
						>
					</div>
				</FormField>
			</div>
		</div>
	</section>

	<!-- EVENT DESCRIPTION -->
	<section>
		<h2 class="mb-4 text-2xl font-semibold">Eventbeskrivelse</h2>
		<div>
			<FormField errors={errors?.fieldErrors.description}>
				<TipTapEditor bind:value={form.description} />
			</FormField>
		</div>
	</section>

	<!-- CONCERTS -->
	<section>
		<h2 class="mb-4 text-2xl font-semibold">Koncerter</h2>
		<FormField errors={errors?.fieldErrors.concerts}>
			<ConcertsList {concerts} {artists} onAdd={addConcert} onDelete={deleteConcert} />
		</FormField>
	</section>

	<Button type="submit" class="w-full md:max-w-64">
		{#if loading}
			<Spinner />
			Offentligører...
		{:else}
			<PublishIcon />
			Offentliggør
		{/if}
	</Button>
</form>
