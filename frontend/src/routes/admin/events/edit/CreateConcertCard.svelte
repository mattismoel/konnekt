<script lang="ts">
	import type { z } from 'zod';

	import type { concertForm } from '$lib/concert';
	import type { Artist } from '$lib/artist';

	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import Selector from '$lib/components/ui/Selector.svelte';
	import DateTimePicker from '$lib/components/ui/DateTimePicker.svelte';

	import PlusIcon from '~icons/mdi/plus';
	import CloseIcon from '~icons/mdi/close';
	import RightArrowIcon from '~icons/mdi/arrow-right';

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
</script>

<Card class="relative flex-1 space-y-4">
	<div class="space-y-8">
		<div class="flex items-center justify-between">
			<h3 class="text-xl font-semibold">#{idx}</h3>
			<button type="button" class="hover:text-text text-zinc-500" onclick={onDelete}
				><CloseIcon /></button
			>
		</div>
		<div class="flex w-full gap-4">
			<Selector
				value={concert.artistID.toString()}
				onchange={(e) => selectArtist(parseInt(e.currentTarget.value))}
				class="w-full"
				entries={artists.map((a) => ({
					name: a.name,
					value: a.id.toString()
				}))}
			/>
			<form action="/admin/artists/edit">
				<Button variant="primary" type="submit">
					<PlusIcon />Ny
				</Button>
			</form>
		</div>
		<div class="flex items-center gap-8">
			<DateTimePicker
				class="w-full"
				label="Fra"
				defaultValue={concert.from}
				onChange={(d) => (concert.from = d)}
			/>
			<DateTimePicker
				class="w-full"
				label="Til"
				defaultValue={concert.to}
				onChange={(d) => (concert.to = d)}
			/>
		</div>
	</div>
</Card>
