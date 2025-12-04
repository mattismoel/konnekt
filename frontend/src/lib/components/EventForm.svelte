<script lang="ts">
	import type { HTMLFormAttributes } from "svelte/elements";
	import FormField from "./FormField.svelte";
	import ImagePreview from "./ImagePreview.svelte";
	import Button from "./ui/Button.svelte";
	import { createEvent, editEvent } from "../features/event.remote.ts";
	import type { Event } from "../features/event";
	import Input from "./ui/Input.svelte";
	import ConcertCard from "./ConcertCard.svelte";
	import { addMinutes, format, roundToNearestHours } from "date-fns";
	import { getArtists } from "../features/artist.remote.ts";
	import { INPUT_DATETIME_FORMAT } from "$lib/time";
	import Select from "./ui/Select.svelte";
	import { getVenues } from "../features/venue.remote.ts";

	const DEFAULT_CHANGEOVER_MINUTES = 15;
	const DEFAULT_CONCERT_DURATION = 40;

	type CreateProps = {
		event?: never;
	};

	type EditProps = {
		event: Event;
	};

	type Props = HTMLFormAttributes & (CreateProps | EditProps);

	let { event, ...rest }: Props = $props();

	const isEditable = true;

	const form = $derived(event ? editEvent : createEvent);

	let coverInput = $state<HTMLInputElement>();

	const addConcert = () => {
		const concertCount = form.fields.concerts.value()?.length ?? 0;

		if (concertCount === 0) {
			const fromDate = roundToNearestHours(new Date());
			const toDate = addMinutes(fromDate, DEFAULT_CONCERT_DURATION);

			form.fields.concerts[0].artistId.set("");
			form.fields.concerts[0].fromDate.set(format(fromDate, INPUT_DATETIME_FORMAT));
			form.fields.concerts[0].toDate.set(format(toDate, INPUT_DATETIME_FORMAT));

			return;
		}

		const lastConcert = form.fields.concerts[concertCount - 1].value();

		const fromDate = addMinutes(lastConcert.toDate, DEFAULT_CHANGEOVER_MINUTES);
		const toDate = addMinutes(fromDate, DEFAULT_CONCERT_DURATION);

		form.fields.concerts[concertCount].artistId.set("");
		form.fields.concerts[concertCount].fromDate.set(format(fromDate, INPUT_DATETIME_FORMAT));
		form.fields.concerts[concertCount].toDate.set(format(toDate, INPUT_DATETIME_FORMAT));
	};
</script>

<form {...form} enctype="multipart/form-data" class="flex flex-col">
	<input {...form.fields.cover.as("file")} />

	<FormField issues={form?.fields.cover.issues()?.map((i) => i.message)} class="mb-16">
		<ImagePreview
			src={form.fields.cover.value()
				? URL.createObjectURL(form.fields.cover.value())
				: event?.cover}
		/>
		<Button type="button" variant="secondary" onclick={() => coverInput?.click()}>Vælg...</Button>
	</FormField>

	<section class="mb-16">
		<h2 class="mb-4 font-medium">Generelt</h2>

		<div class="flex flex-col gap-4">
			<FormField issues={form.fields.title.issues()?.map((i) => i.message)}>
				<Input {...form.fields.title.as("text")} placeholder="Eventtitel" />
			</FormField>

			<FormField issues={form.fields.description.issues()?.map((i) => i.message)}>
				<Input {...form.fields.description.as("text")} placeholder="Beskrivelse" />
			</FormField>

			<fieldset class="flex gap-2">
				<FormField issues={form.fields.ticketUrl.issues()?.map((i) => i.message)}>
					<Input {...form.fields.ticketUrl.as("url")} placeholder="Billet-URL" />
				</FormField>

				{#await getVenues(undefined) then { items: venues }}
					<FormField issues={form.fields.title.issues()?.map((i) => i.message)}>
						<Select {...form.fields.venueId.as("select")}>
							{#each venues as venue}
								<option value={venue.id}>{venue.name}</option>
							{/each}
						</Select>
					</FormField>
				{/await}
				<Button href="/admin/venues/create" target="_blank">+</Button>
			</fieldset>

			<label for="" class="flex items-center gap-4">
				<input {...form.fields.isPublic.as("checkbox")} />
				Offentlig
			</label>
		</div>
	</section>

	{#await getArtists(undefined) then { items: artists }}
		<section class="mb-16">
			<h2 class="mb-4 font-medium">Koncerter</h2>

			<ul class="flex flex-col gap-2">
				{#each form.fields.concerts.value(), i}
					<ConcertCard isEdit={event !== undefined} {artists} idx={i} />
				{/each}
				<Button type="button" variant="secondary" onclick={addConcert}>Tilføj</Button>
			</ul>
		</section>
	{/await}

	<FormField issues={[]}>
		<!-- <Controller -->
		<!--   control={control} -->
		<!--   name="description" -->
		<!--   render={({ field: { value, onChange } }) => ( -->
		<!--     <Tiptap -->
		<!--       disabled={!isEditable} -->
		<!--       content={value} -->
		<!--       onChange={onChange} -->
		<!--     /> -->
		<!--   )} -->
		<!-- /> -->
	</FormField>

	<p>
		{form.fields
			.allIssues()
			?.map((i) => `${i.path}, ${i.message}`)
			.join(", ")}
	</p>

	<div class="flex flex-col gap-4">
		{#if isEditable}
			<Button class="w-full" type="submit">
				{#if form.fields.isPublic.value() === true}
					Offentligør
				{:else}
					Upload
				{/if}
			</Button>

			{#if event}
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

<!-- const GeneralSection = () => { -->
<!--   const { -->
<!--     register, -->
<!--     formState: { errors, disabled }, -->
<!--   } = useEventFormContext(); -->
<!--   const isEditable = !disabled; -->
<!---->
<!--   return ( -->
<!--     <section> -->
<!--       <h1 class="mb-4 font-heading text-2xl font-bold">Generelt</h1> -->
<!--       <div class="flex flex-col gap-4"> -->
<!--         <FormField error={errors.title}> -->
<!--           <Input {...register("title")} placeholder="Eventtitel" /> -->
<!--         </FormField> -->
<!---->
<!--         <div class="flex flex-col gap-12 @xl:flex-row"> -->
<!--           <FormField error={errors.ticketUrl}> -->
<!--             <Input -->
<!--               {...register("ticketUrl")} -->
<!--               placeholder="Billet-URL" -->
<!--               class="w-full" -->
<!--             /> -->
<!--           </FormField> -->
<!---->
<!--           <VenueSelector /> -->
<!--         </div> -->
<!--         {isEditable && ( -->
<!--           <FormField class="w-min"> -->
<!--             <label class="flex items-center gap-2"> -->
<!--               <input type="checkbox" {...register("isPublic")} /> -->
<!--               Offentlig -->
<!--             </label> -->
<!--           </FormField> -->
<!--         )} -->
<!--       </div> -->
<!--     </section> -->
<!--   ); -->
<!-- }; -->
<!---->
<!-- const ConcertsSection = () => { -->
<!--   const { fields } = useEventFormContext(); -->
<!--   return ( -->
<!--     <section> -->
<!--       <h1 class="mb-4 font-heading text-2xl font-bold">Koncerter</h1> -->
<!--       <ConcertList> -->
<!--         {fields.map((field, index) => ( -->
<!--           <ConcertList.Entry key={field.id} index={index} /> -->
<!--         ))} -->
<!--       </ConcertList> -->
<!--     </section> -->
<!--   ); -->
<!-- }; -->
<!---->
<!-- const VenueSelector = () => { -->
<!--   const { -->
<!--     venues, -->
<!--     control, -->
<!--     formState: { disabled }, -->
<!--   } = useEventFormContext(); -->
<!--   const isEditable = !disabled; -->
<!---->
<!--   return ( -->
<!--     <Controller -->
<!--       control={control} -->
<!--       name="venueId" -->
<!--       render={({ field: { onChange, ...rest }, fieldState: { error } }) => ( -->
<!--         <FormField error={error}> -->
<!--           <Selector -->
<!--             {...rest} -->
<!--             onChange={(e) => onChange(parseInt(e.target.value))} -->
<!--             placeholder="Vælg venue..." -->
<!--             class="h-min w-full" -->
<!--           > -->
<!--             {venues.map(({ id, name }) => ( -->
<!--               <option key={id} value={id}> -->
<!--                 {name} -->
<!--               </option> -->
<!--             ))} -->
<!--           </Selector> -->
<!---->
<!--           {isEditable && ( -->
<!--             <div class="flex gap-2"> -->
<!--               <Button variant="secondary" class="h-full"> -->
<!--                 <FaArrowsRotate /> -->
<!--               </Button> -->
<!--               <LinkButton -->
<!--                 to="/admin/venues/create" -->
<!--                 class="h-full" -->
<!--                 target="__blank" -->
<!--               > -->
<!--                 <FaPlus /> -->
<!--               </LinkButton> -->
<!--             </div> -->
<!--           )} -->
<!--         </FormField> -->
<!--       )} -->
<!--     /> -->
<!--   ); -->
<!-- }; -->
<!---->
<!-- export default EventForm; -->
