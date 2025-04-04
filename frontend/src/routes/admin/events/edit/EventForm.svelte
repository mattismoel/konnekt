<script lang="ts">
	import { z, ZodError } from 'zod';
	import { addMinutes, roundToNearestHours } from 'date-fns';

	import type { Artist } from '$lib/artist';
	import type { Venue } from '$lib/venue';
	import { eventForm, type Event } from '$lib/event';
	import { concertForm } from '$lib/concert';

	import Button from '$lib/components/ui/Button.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import FieldError from '$lib/components/ui/FieldError.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import ImagePreview from '$lib/components/ui/ImagePreview.svelte';
	import CreateConcertCard from './CreateConcertCard.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import { goto } from '$app/navigation';
	import TipTapEditor from '$lib/components/ui/tiptap/TipTapEditor.svelte';

	type Props = {
		event: Event | null;
		artists: Artist[];
		venues: Venue[];

		editable?: boolean;

		onSubmit: (e: z.infer<typeof eventForm>) => void;
	};

	const { event, artists, venues, editable = false, onSubmit }: Props = $props();

	const concertWithID = concertForm.extend({ id: z.string().uuid() });

	const extendedForm = eventForm.extend({
		concerts: concertWithID.array()
	});

	let form = $state<z.infer<typeof extendedForm>>({
		title: event?.title || '',
		description: event?.description || '',
		ticketUrl: event?.ticketUrl || '',
		venueId: event?.venue.id || 1,
		image: null,
		concerts:
			event?.concerts.map((c) => ({
				id: crypto.randomUUID().toString(),
				artistID: c.artist.id,
				from: c.from,
				to: c.to
			})) || []
	});

	let formError = $state<ZodError>();
	let imageUrl = $derived(form.image ? URL.createObjectURL(form.image) : event?.imageUrl || '');

	const deleteConcert = (id: string) => {
		form.concerts = form.concerts.filter((c) => c.id !== id);
	};

	const addConcert = () => {
		const from = form.concerts.length > 0 ? form.concerts[0]?.to : roundToNearestHours(new Date());
		const to = addMinutes(from, 30);

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

<form class="space-y-16" onsubmit={submit}>
	<h1 class="font-heading mb-8 text-4xl font-bold">
		{#if event}
			Redigér event
		{:else}
			Lav event
		{/if}
	</h1>
	<!-- GENERAL -->
	<div>
		<div>
			<ImagePreview src={imageUrl || ''} onChange={(file) => (form.image = file)} />
			<FieldError errors={formError?.flatten().fieldErrors['imageUrl']} />
		</div>
		<FieldError errors={formError?.flatten().fieldErrors['image']} />
		<div class="space-y-8">
			<div class="flex gap-4 *:flex-1">
				<Input
					errors={formError?.flatten().fieldErrors['title']}
					label="Titel"
					type="text"
					name="title"
					bind:value={form.title}
				/>
				<div class="flex flex-col">
					<div class="flex gap-2">
						<Selector
							class="w-full"
							onchange={(e) => (form.venueId = parseInt(e.currentTarget.value))}
							value={form.venueId.toString()}
							entries={venues.map((v) => ({
								name: `${v.name}, ${v.city}`,
								value: v.id.toString()
							}))}
						/>
						<Button disabled={editable} type="submit" onclick={() => goto('/admin/venues')}
							><PlusIcon />Tilføj</Button
						>
					</div>
					<FieldError errors={formError?.flatten().fieldErrors['venueId']} />
				</div>
			</div>
			<Input
				errors={formError?.flatten().fieldErrors['ticketUrl']}
				label="Billet-URL"
				bind:value={form.ticketUrl}
			/>
		</div>
	</div>

	<!-- EVENT DESCRIPTION -->
	<div>
		<h1 class="font-heading mb-8 text-4xl font-bold">Eventbeskrivelse</h1>
		<TipTapEditor bind:value={form.description} />
	</div>

	<!-- CONCERTS SECTION -->
	<div class="space-y-6">
		<h1 class="font-heading mb-8 text-4xl font-bold">Koncerter</h1>
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
		<Button onclick={addConcert} variant="ghost">+ Tilføj koncert</Button>
		<FieldError errors={formError?.flatten().fieldErrors['concerts']} />
	</div>

	<!-- ACTIONS -->
	<div class="flex gap-4">
		<Button variant="secondary">Preview</Button>
		<Button type="submit">Offentligør</Button>
	</div>
</form>
