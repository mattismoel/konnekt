<script lang="ts">
	import type { Artist } from '$lib/artist';
	import Button from '$lib/components/Button.svelte';
	import DateTimePicker from '$lib/components/DateTimePicker.svelte';
	import Selector from '$lib/components/Selector.svelte';
	import PlusIcon from '~icons/mdi/plus';
	import RefreshIcon from '~icons/mdi/refresh';
	import CloseIcon from '~icons/mdi/close';
	import RightArrowIcon from '~icons/mdi/arrow-right';
	import type { UpdateConcert } from '$lib/concert';
	import { goto, invalidateAll } from '$app/navigation';
	import Card from '$lib/components/Card.svelte';

	type Props = {
		artists: Artist[];
		from?: Date;
		to?: Date;
		idx: number;

		concert: UpdateConcert;

		onDelete: () => void;
	};

	let { artists, concert = $bindable(), from, to, idx, onDelete }: Props = $props();

	const selectArtist = (artistId: number) => {
		concert.artistID = artistId;
	};

	$inspect(artists);
</script>

<Card class="relative flex-1 space-y-4">
	<button
		type="button"
		class="hover:text-text absolute top-4 right-4 text-zinc-500"
		onclick={onDelete}><CloseIcon /></button
	>
	<div class="space-y-8">
		<h3 class="text-xl font-semibold">#{idx}</h3>
		<div class="flex w-full gap-4">
			<Selector
				onchange={(e) => selectArtist(parseInt(e.currentTarget.value))}
				class="w-full"
				entries={artists.map((a) => ({
					name: a.name,
					value: a.id.toString()
				}))}
			/>
			<button type="button" onclick={invalidateAll}>
				<RefreshIcon class="text-zinc-500" />
			</button>
			<Button variant="primary" onclick={() => goto('/admin/artists/edit')}>
				<PlusIcon />Ny
			</Button>
		</div>
		<div class="flex items-center gap-8">
			<DateTimePicker label="Fra" bind:date={concert.from} />
			<RightArrowIcon />
			<DateTimePicker label="Til" bind:date={concert.to} />
		</div>
	</div>
</Card>
