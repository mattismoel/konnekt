<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import FormField from '$lib/components/ui/FormField.svelte';
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

<form onsubmit={(e) => handleSubmit(e)} class="flex max-w-lg flex-col gap-8">
	<h1 class="text-4xl font-bold">
		{#if venue}
			Redigér venue
		{:else}
			Lav venue
		{/if}
	</h1>

	<div class="flex flex-col gap-4">
		<FormField errors={errors?.fieldErrors.name}>
			<Input placeholder="Venuenavn" bind:value={form.name} />
		</FormField>

		<div class="flex gap-4">
			<FormField errors={errors?.fieldErrors.city}>
				<Input placeholder="By" class="flex-1" bind:value={form.city} />
			</FormField>
			<div>
				<FormField errors={errors?.fieldErrors.countryCode}>
					<Selector
						class="h-min w-min"
						entries={Array.from(COUNTRIES_MAP).map(([key, val]) => ({ name: val, value: key }))}
						bind:value={form.countryCode}
					/>
				</FormField>
			</div>
		</div>
	</div>
	<Button type="submit" class="w-full">Offentliggør</Button>
</form>
