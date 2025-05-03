<script lang="ts">
	import type { z } from 'zod';

	import type { concertForm } from '$lib/features/concert/concert';
	import type { Artist } from '$lib/features/artist/artist';

	import Button from '$lib/components/ui/Button.svelte';
	import * as Card from '$lib/components/ui/card/index';
	import Selector from '$lib/components/ui/Selector.svelte';
	import DateTimePicker from '$lib/components/DateTimePicker.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import CloseIcon from '~icons/mdi/close';
	import { goto } from '$app/navigation';

	type Props = {
		artists: Artist[];
		idx: number;

		concert: z.infer<typeof concertForm>;

		onDelete: () => void;
	};

	let { artists, concert = $bindable(), idx, onDelete }: Props = $props();

	const selectArtist = (artistId: number) => {
		concert.artistID = artistId;
	};

	const handleDelete = () => {
		const artist = artists.find((a) => a.id === concert.artistID);

		if (!artist) return;

		if (!confirm(`Er du sikker på, at du vil slette koncerten med ${artist.name}?`)) {
			return;
		}

		onDelete();
	};
</script>

<Card.Root class="relative flex flex-1 flex-col">
	<Card.Header>
		<Card.Title>#{idx + 1}</Card.Title>
		<button
			type="button"
			class="text-text/50 hover:text-text absolute top-6 right-6"
			onclick={handleDelete}><CloseIcon /></button
		>
	</Card.Header>

	<Card.Content class="gap-8">
		<div class="flex w-full gap-4">
			<Selector
				value={concert.artistID.toString()}
				onchange={(e) => selectArtist(parseInt(e.currentTarget.value))}
				class="w-full"
				entries={artists.map((a) => ({
					name: a.name,
					value: a.id.toString()
				}))}
			></Selector>
			<Button href="/admin/artists/create" title="Tilføj ny kunstner" variant="ghost">
				<PlusIcon></PlusIcon>Lav&nbsp;ny
			</Button>
		</div>
		<div class="flex flex-col items-center gap-4 sm:flex-row">
			<DateTimePicker
				class="w-full"
				placeholder="Fra"
				defaultValue={concert.from}
				onChange={(d) => (concert.from = d)}
			/>
			<DateTimePicker
				class="w-full"
				placeholder="Til"
				defaultValue={concert.to}
				onChange={(d) => (concert.to = d)}
			/>
		</div>
	</Card.Content>
</Card.Root>
