<script lang="ts">
	import { concertForm } from '$lib/features/concert/concert';
	import { z } from 'zod';
	import ConcertCard from './ConcertCard.svelte';
	import type { Artist } from '$lib/features/artist/artist';
	import Button from '$lib/components/ui/Button.svelte';
	import AddIcon from '~icons/mdi/plus';

	type Props = {
		artists: Artist[];
		concerts: Map<string, z.infer<typeof concertForm>>;

		onAdd: () => void;
		onDelete: (id: string) => void;
	};

	let { concerts, artists, onAdd, onDelete }: Props = $props();
</script>

<div class="flex flex-col gap-4">
	{#each Array.from(concerts.entries()) as [id, concert], idx (id)}
		<ConcertCard {concert} {artists} {idx} onDelete={() => onDelete(id)} />
	{/each}
</div>
<Button variant="ghost" onclick={onAdd}><AddIcon />Tilføj</Button>
