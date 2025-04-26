<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import FieldError from '$lib/components/ui/FieldError.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import { COUNTRIES_MAP } from '$lib/features/venue/countries';
	import type { createVenueForm, editVenueForm, Venue } from '$lib/features/venue/venue';
	import type { z } from 'zod';

	type Props = {
		venue?: Venue;

		errors:
			| z.typeToFlattenedError<z.infer<typeof createVenueForm>>
			| z.typeToFlattenedError<z.infer<typeof editVenueForm>>
			| undefined;
		onSubmit: (form: z.infer<typeof createVenueForm> | z.infer<typeof editVenueForm>) => void;
	};

	let { venue, errors, onSubmit }: Props = $props();

	const form = $state<z.infer<typeof createVenueForm> | z.infer<typeof editVenueForm>>(
		venue || {
			name: '',
			city: 'Odense',
			countryCode: 'DK'
		}
	);

	const handleSubmit = (e: SubmitEvent) => {
		e.preventDefault();
		onSubmit(form);
	};
</script>

<form onsubmit={(e) => handleSubmit(e)} class="space-y-8">
	<h1 class="mb-8 text-4xl font-bold">
		{#if venue}
			Lav venue
		{:else}
			Redigér venue
		{/if}
	</h1>

	<div class="flex flex-col gap-4">
		<Input label="Navn" bind:value={form.name} errors={errors?.fieldErrors.name} />
		<div class="flex gap-4">
			<Input label="By" class="flex-1" bind:value={form.city} errors={errors?.fieldErrors.city} />
			<div>
				<Selector
					class="h-min w-min"
					entries={Array.from(COUNTRIES_MAP).map(([key, val]) => ({ name: val, value: key }))}
					bind:value={form.countryCode}
				/>
				<FieldError errors={errors?.fieldErrors.countryCode} />
			</div>
		</div>
	</div>
	<Button type="submit" class="w-full">Offentliggør</Button>
</form>
