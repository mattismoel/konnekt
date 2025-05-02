<script lang="ts">
	import type { Team } from '$lib/features/auth/team';
	import * as Modal from '$lib/components/ui/modal/index';
	import SelectorEntry from './ui/SelectorEntry.svelte';
	import Button from './ui/Button.svelte';
	import { Description } from './ui/card';
	import SearchBar from './SearchBar.svelte';

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

<Modal.Root {show}>
	<Modal.Header {onClose}>
		<Modal.Title>Vælg medlemsteams...</Modal.Title>
		<Modal.Description
			>Her kan du vælge de teams, som medlemmet er del af, og dermed hvilke rettigheder det har.</Modal.Description
		>
	</Modal.Header>
	<Modal.Content class="flex flex-col gap-4">
		<SearchBar bind:value={search} />
		<div class="flex flex-col gap-2">
			{#each teams as team (team.id)}
				<SelectorEntry
					selected={selected.some((t) => t.id === team.id)}
					name={team.displayName}
					onSelect={() => select(team.id)}
					onDeselect={() => deselect(team.id)}
				/>
			{/each}
		</div>
	</Modal.Content>
	<Modal.Footer>
		<Button onclick={onClose}>Vælg</Button>
	</Modal.Footer>
</Modal.Root>
