<script lang="ts">
	import PlusIcon from '~icons/mdi/add';
	import SelectorEntry from './ui/SelectorEntry.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import { error } from '@sveltejs/kit';
	import { invalidateAll } from '$app/navigation';
	import { createGenre, type Genre } from '$lib/features/artist/genre';
	import * as Modal from '$lib/components/ui/modal/index';
	import { APIError } from '$lib/api';

	type Props = {
		genres: Genre[];
		show: boolean;
		onClose: () => void;
		onChange: (selected: Genre[]) => void;
	};

	let { genres, show, onClose, onChange }: Props = $props();

	let search = $state('');
	let selected = $state<Genre[]>([]);

	const addGenre = async () => {
		try {
			await createGenre(fetch, search);
			await invalidateAll();
		} catch (e) {
			if (e instanceof APIError) return error(e.status, e.message);
			return error(500, 'Could not create genre' + e);
		}
	};

	const select = (genre: Genre) => {
		selected = [...selected, genre];
		onChange(selected);
	};

	const deselect = (genre: Genre) => {
		selected = selected.filter((g) => g.id !== genre.id);
		onChange(selected);
	};
</script>

<Modal.Root {show}>
	<Modal.Header {onClose}>
		<Modal.Title>Vælg genrer...</Modal.Title>
		<Modal.Description>Her kan du vælge de genrer, som kunstneren associeres med.</Modal.Description
		>
	</Modal.Header>

	<Modal.Content class="text-text space-y-4">
		<div class="flex gap-2">
			<input
				type="text"
				placeholder="Søg..."
				bind:value={search}
				class="rounded-md border border-zinc-800 bg-zinc-900"
			/>
			<Button type="button" onclick={addGenre}>
				<PlusIcon />Tilføj
			</Button>
		</div>
		<div class="space-y-1">
			{#each genres as genre (genre.id)}
				<SelectorEntry
					selected={selected.some((g) => g.id === genre.id)}
					name={genre.name}
					onSelect={() => select(genre)}
					onDeselect={() => deselect(genre)}
				/>
			{/each}
		</div>
	</Modal.Content>
	<Modal.Footer>
		<Button type="button" onclick={onClose}>Vælg</Button>
	</Modal.Footer>
</Modal.Root>
