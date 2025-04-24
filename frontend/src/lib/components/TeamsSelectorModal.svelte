<script lang="ts">
	import type { Team } from '$lib/features/auth/team';
	import Modal from './ui/modal/Modal.svelte';
	import ModalContent from './ui/modal/ModalContent.svelte';
	import ModalFooter from './ui/modal/ModalFooter.svelte';
	import ModalHeader from './ui/modal/ModalHeader.svelte';
	import SelectorEntry from './ui/SelectorEntry.svelte';

	type Props = {
		teams: Team[];

		selected: Team[];

		show: boolean;
		onClose: () => void;
		onChange: (selected: Team[]) => void;
	};

	let { teams, selected = $bindable([]), show, onChange, onClose }: Props = $props();
	let search = $state('');

	const select = (id: number) => {
		const team = teams.find((team) => team.id === id);
		if (!team) return;

		selected = [...selected, team];
		onChange(selected);
	};

	const deselect = (id: number) => {
		selected = selected.filter((team) => team.id !== id);
		onChange(selected);
	};
</script>

<Modal {show}>
	<ModalHeader label="Vælg medlemsteams..." {onClose} />
	<ModalContent class="text-text space-y-4">
		<div class="flex gap-2">
			<input
				type="text"
				placeholder="Søg..."
				bind:value={search}
				class="rounded-md border border-zinc-800 bg-zinc-900"
			/>
			<!-- <Button type="button" onclick={addGenre}> -->
			<!-- 	<AddIcon />Tilføj -->
			<!-- </Button> -->
		</div>
		<div class="space-y-1">
			{#each teams as team (team.id)}
				<SelectorEntry
					selected={selected.some((t) => t.id === team.id)}
					name={team.displayName}
					onSelect={() => select(team.id)}
					onDeselect={() => deselect(team.id)}
				/>
			{/each}
		</div>
	</ModalContent>
	<ModalFooter>Hey</ModalFooter>
</Modal>
