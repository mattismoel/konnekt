<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Pill from '$lib/components/Pill.svelte';
	import TeamsSelectorModal from '$lib/components/TeamsSelectorModal.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import FieldError from '$lib/components/ui/FieldError.svelte';
	import { setMemberTeams, type Member, type setMemberTeamsForm } from '$lib/features/auth/member';
	import type { Team } from '$lib/features/auth/team';
	import { toaster } from '$lib/toaster.svelte';
	import { ZodError, type z } from 'zod';
	import EditIcon from '~icons/mdi/edit';

	let isModalOpen = $state(false);

	type Props = {
		member: Member;
		defaultSelected?: Team[];
		teams: Team[];
	};

	let { member, teams, defaultSelected }: Props = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof setMemberTeamsForm>>>();

	let selected = $state<Team[]>(defaultSelected || []);

	let hasChanged = $derived.by(() => {
		if (selected.length !== defaultSelected?.length) return true;

		return selected.some(
			(selectedTeam) => !defaultSelected?.find((defaultTeam) => defaultTeam.id === selectedTeam.id)
		);
	});

	const handleSubmit = async (e: SubmitEvent) => {
		e.preventDefault();

		try {
			// This is okay, as we've guarded selected teams array length up above!
			await setMemberTeams(
				fetch,
				member.id,
				selected.map((team) => team.id)
			);
			await invalidateAll();
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
			}

			toaster.addToast('Kunne ikke opdatere medlemmets hold', 'Noget gik galt...', 'error');
		}
	};
</script>

<h1 class="mb-4 text-2xl font-bold">Teams</h1>
<form onsubmit={handleSubmit} class="flex flex-wrap gap-4">
	<ul class="flex flex-wrap gap-2">
		<button
			type="button"
			onclick={() => (isModalOpen = true)}
			class="flex items-center gap-2 rounded-full border border-zinc-900 px-4 py-2 hover:bg-zinc-900"
			><EditIcon />Redigér</button
		>
		{#each selected || [] as team (team.id)}
			<Pill>{team.displayName}</Pill>
		{/each}
	</ul>
	<FieldError errors={errors?.formErrors} />
	<Button disabled={!hasChanged} type="submit" class="w-full">Opdatér</Button>

	<TeamsSelectorModal
		{teams}
		selected={selected || []}
		show={isModalOpen}
		onClose={() => (isModalOpen = false)}
		onChange={(newTeams) => (selected = newTeams)}
	/>
</form>
