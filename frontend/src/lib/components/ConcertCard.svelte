<script lang="ts">
	import type { Artist } from "$lib/features/artist";
	import { createEvent, editEvent } from "$lib/features/event.remote";
	import Card from "./Card.svelte";
	import FormField from "./FormField.svelte";
	import Input from "./ui/Input.svelte";
	import Select from "./ui/Select.svelte";

	type Props = {
		idx: number;
		artists: Artist[];
		isEdit: boolean;
	};

	let { artists, idx, isEdit }: Props = $props();

	const form = $derived(isEdit ? createEvent : editEvent);
</script>

<li>
	<Card title="#{idx + 1}">
		<FormField
			issues={form.fields.concerts[idx].artistId.issues()?.map((i) => i.message)}
			class="mb-4"
		>
			<Select class="w-full" {...form.fields.concerts[idx].artistId.as("select")}>
				{#each artists as artist}
					<option selected value={artist.id}>{artist.name}</option>
				{/each}
			</Select>
		</FormField>

		<fieldset class="flex items-center gap-2">
			<FormField issues={form.fields.concerts[idx].fromDate.issues()?.map((i) => i.message)}>
				<Input {...form.fields.concerts[idx].fromDate.as("datetime-local")} />
			</FormField>

			<p>&gt;</p>

			<FormField issues={form.fields.concerts[idx].toDate.issues()?.map((i) => i.message)}>
				<Input {...form.fields.concerts[idx].toDate.as("datetime-local")} />
			</FormField>
		</fieldset>
	</Card>
</li>
