<script lang="ts">
	import type { HTMLFormAttributes } from "svelte/elements";
	import FormField from "$lib/components/FormField.svelte";
	import ImagePreview from "$lib/components/ImagePreview.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import { createEvent, editEvent } from "../event.remote.ts";
	import type { Event } from "../event";
	import Input from "$lib/components/ui/Input.svelte";
	import { addMinutes, format, roundToNearestHours } from "date-fns";
	import { getArtists } from "../../artists/artist.remote.ts";
	import { INPUT_DATETIME_FORMAT } from "$lib/time";
	import Select from "$lib/components/ui/Select.svelte";
	import { getVenues } from "../../venues/venue.remote.ts";
	import Card from "$lib/components/Card.svelte";
	import { getToastContext } from "../../../components/toaster/toaster.svelte.ts";
	import { goto } from "$app/navigation";

	const DEFAULT_CHANGEOVER_MINUTES = 15;
	const DEFAULT_CONCERT_DURATION = 40;

	type CreateProps = {
		variant: "create";
		form: typeof createEvent;
		event?: never;
	};

	type EditProps = {
		variant: "edit";
		form: typeof editEvent;
		event: Event;
	};

	type Props = HTMLFormAttributes & (CreateProps | EditProps);

	let { ...rest }: Props = $props();

	let { title, description, cover, ticketUrl, venueId, isPublic } = rest.form.fields;

	const isEditable = true;

	const toaster = getToastContext();

	let coverInput = $state<HTMLInputElement>();

	$effect(() => {
		if (rest.variant === "create") return;

		rest.form.fields.set({
			...rest.event,
			eventId: rest.event.id,
			cover: undefined,
			venueId: rest.event.venue.id,
			concerts: rest.event.concerts.map((concert) => ({
				id: concert.id,
				artistId: concert.artist.id,
				fromDate: format(concert.fromDate, INPUT_DATETIME_FORMAT),
				toDate: format(concert.toDate, INPUT_DATETIME_FORMAT)
			}))
		});
	});

	const addConcert = () => {
		const concertCount = rest.form.fields.concerts.value()?.length ?? 0;

		if (concertCount === 0) {
			console.log("first concert");
			const fromDate = roundToNearestHours(new Date());
			const toDate = addMinutes(fromDate, DEFAULT_CONCERT_DURATION);

			rest.form.fields.concerts[0].set({
				id: undefined,
				artistId: "",
				fromDate: format(fromDate, INPUT_DATETIME_FORMAT),
				toDate: format(toDate, INPUT_DATETIME_FORMAT)
			});

			return;
		}

		const lastConcert = rest.form.fields.concerts[concertCount - 1].value();

		const fromDate = addMinutes(lastConcert.toDate, DEFAULT_CHANGEOVER_MINUTES);
		const toDate = addMinutes(fromDate, DEFAULT_CONCERT_DURATION);

		rest.form.fields.concerts[concertCount].set({
			id: undefined,
			artistId: "",
			fromDate: format(fromDate, INPUT_DATETIME_FORMAT),
			toDate: format(toDate, INPUT_DATETIME_FORMAT)
		});
	};
</script>

<form
	{...rest.form.enhance(async ({ submit }) => {
		try {
			await submit();
			toaster.add(rest.variant === "create" ? "Event tilføjet" : "Event opdateret");
			goto("/admin/events");
		} catch (e) {
			toaster.add(
				rest.variant === "create" ? "Kunne ikke lave event" : "Kunne ikke opdatere event"
			);
			console.error(e);
			throw e;
		}
	})}
	enctype="multipart/form-data"
	class="flex flex-col"
>
	{#if rest.variant === "edit"}
		<input {...rest.form.fields.eventId.as("hidden", rest.event.id)} />
	{/if}

	<FormField issues={cover.issues()?.map((i) => i.message)} class="mb-16">
		<input bind:this={coverInput} {...rest.form.fields.cover.as("file")} class="hidden" />
		<ImagePreview src={cover.value() ? URL.createObjectURL(cover.value()) : rest.event?.cover} />
		<Button type="button" variant="secondary" onclick={() => coverInput?.click()}>Vælg...</Button>
	</FormField>

	<section class="mb-16">
		<h2 class="mb-4 font-medium">Generelt</h2>

		<div class="flex flex-col gap-4">
			<FormField issues={title.issues()?.map((i) => i.message)}>
				<Input {...title.as("text")} placeholder="Eventtitel" />
			</FormField>

			<FormField issues={description.issues()?.map((i) => i.message)}>
				<Input {...description.as("text")} placeholder="Beskrivelse" />
			</FormField>

			<fieldset class="flex gap-2">
				<FormField issues={ticketUrl.issues()?.map((i) => i.message)}>
					<Input {...ticketUrl.as("url")} placeholder="Billet-URL" />
				</FormField>

				{#await getVenues(undefined) then { items: venues }}
					<FormField issues={title.issues()?.map((i) => i.message)}>
						<Select {...venueId.as("select")}>
							{#each venues as venue}
								<option value={venue.id}>{venue.name}</option>
							{/each}
						</Select>
					</FormField>
				{/await}
				<Button href="/admin/venues/create" target="_blank">+</Button>
			</fieldset>

			<label for="" class="flex items-center gap-4">
				<input {...isPublic.as("checkbox")} />
				Offentlig
			</label>
		</div>
	</section>

	{#await getArtists(undefined) then { items: artists }}
		<section class="mb-16">
			<h2 class="mb-4 font-medium">Koncerter</h2>

			<ul class="flex flex-col gap-2">
				{#each rest.form.fields.concerts.value(), i}
					<li>
						<Card title="#{i + 1}">
							{#if rest.variant === "edit" && rest.form.fields.concerts[i].id.value()}
								<input
									{...rest.form.fields.concerts[i].id.as(
										"hidden",
										rest.form.fields.concerts[i].id.value()
									)}
								/>
							{/if}

							<div class="flex gap-2">
								<FormField
									issues={rest.form.fields.concerts[i].artistId.issues()?.map((i) => i.message)}
									class="mb-4"
								>
									<Select class="w-full" {...rest.form.fields.concerts[i].artistId.as("select")}>
										{#each artists as artist}
											<option selected value={artist.id}>{artist.name}</option>
										{/each}
									</Select>
								</FormField>

								<Button href="/admin/artists/create" target="_blank" class="h-min">+</Button>

								<Button
									type="button"
									variant="secondary"
									class="h-min"
									onclick={getArtists(undefined).refresh}>Refresh</Button
								>
							</div>

							<fieldset class="flex items-center gap-2">
								<FormField
									issues={rest.form.fields.concerts[i].fromDate.issues()?.map((i) => i.message)}
								>
									<Input {...rest.form.fields.concerts[i].fromDate.as("datetime-local")} />
								</FormField>

								<p>&gt;</p>

								<FormField
									issues={rest.form.fields.concerts[i].toDate.issues()?.map((i) => i.message)}
								>
									<Input {...rest.form.fields.concerts[i].toDate.as("datetime-local")} />
								</FormField>
							</fieldset>
						</Card>
					</li>
				{/each}
				<Button type="button" variant="secondary" onclick={addConcert}>Tilføj</Button>
			</ul>
		</section>
	{/await}

	<p>
		{rest.form.fields
			.allIssues()
			?.map((i) => `${i.path}, ${i.message}`)
			.join(", ")}
	</p>

	<div class="flex flex-col gap-4">
		{#if isEditable}
			<Button class="w-full" type="submit">
				{#if isPublic.value() === true}
					Offentligør
				{:else}
					Upload
				{/if}
			</Button>

			{#if rest.event}
				<Button
					type="button"
					variant="dangerous"
					class="w-full"
					onclick={() => alert("Not implemented...")}
				>
					Slet
				</Button>
			{/if}
		{/if}
	</div>
</form>
