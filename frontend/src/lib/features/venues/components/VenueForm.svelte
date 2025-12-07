<script lang="ts">
	import FormField from "$lib/components/FormField.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import Select from "$lib/components/ui/Select.svelte";
	import { hasPermissions } from "$lib/features/auth/auth.remote";
	import { COUNTRIES_MAP } from "$lib/location";
	import type { Venue } from "../venue";
	import { type createVenue, type editVenue } from "../venue.remote";

	type CreateProps = {
		variant: "create";
		form: typeof createVenue;
		venue?: never;
	};

	type EditProps = {
		variant: "edit";
		form: typeof editVenue;
		venue: Venue;
	};

	type Props = CreateProps | EditProps;

	let { ...rest }: Props = $props();

	$effect(() => {
		if (rest.variant === "create") {
			rest.form.fields.countryCode.set("DK");
			return;
		}

		rest.form.fields.set({
			...rest.venue,
			venueId: rest.venue.id
		});
	});

	const disabled = $derived(!(await hasPermissions(["venues:edit"])));
</script>

<form {...rest.form}>
	<h1 class="font-heading mb-8 text-4xl font-bold text-text-light">
		{#if rest.variant === "create"}
			Nyt venue
		{:else}
			Redigér venue
		{/if}
	</h1>

	<div class="mb-8 flex flex-col gap-2">
		{#if rest.variant === "edit"}
			<input {...rest.form.fields.venueId.as("hidden", rest.venue.id)} />
		{/if}
		<FormField issues={rest.form.fields.name.issues()?.map((i) => i.message)}>
			<Input {...rest.form.fields.name.as("text")} placeholder="Venuenavn" {disabled} />
		</FormField>

		<div class="flex gap-2">
			<FormField issues={rest.form.fields.city.issues()?.map((i) => i.message)}>
				<Input {...rest.form.fields.city.as("text")} placeholder="By" {disabled} />
			</FormField>

			<FormField issues={rest.form.fields.countryCode.issues()?.map((i) => i.message)}>
				<Select {...rest.form.fields.countryCode.as("select")} {disabled}>
					{#each COUNTRIES_MAP as [code, name]}
						<option value={code}>
							{name}
						</option>
					{/each}
				</Select>
			</FormField>
		</div>
	</div>

	{#if !disabled}
		<div class="flex flex-col gap-2">
			{#if rest.variant === "edit"}
				<Button variant="dangerous">Slet</Button>
			{/if}

			<Button class="w-full">
				{#if rest.variant === "create"}
					Lav venue
				{:else}
					Redigér venue
				{/if}
			</Button>
		</div>
	{/if}
</form>
