<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Pill from '$lib/components/Pill.svelte';
	import TeamsSelectorModal from '$lib/components/TeamsSelectorModal.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import FormField from '$lib/components/ui/FormField.svelte';
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

<div class="flex flex-col gap-8">
	<h1 class="text-2xl font-bold">Teams</h1>

	<form onsubmit={handleSubmit} class="flex flex-wrap gap-8">
		<FormField errors={errors?.formErrors}>
			<ul class="flex flex-wrap gap-2">
				<Button
					variant="ghost"
					type="button"
					onclick={() => (isModalOpen = true)}
					class="rounded-full"><EditIcon />Redigér</Button
				>
				{#each selected || [] as team (team.id)}
					<Pill>{team.displayName}</Pill>
				{/each}
			</ul>
		</FormField>
		<Button disabled={!hasChanged} type="submit" class="w-full">Opdatér</Button>

		<TeamsSelectorModal
			{teams}
			selected={selected || []}
			bind:show={isModalOpen}
			onChange={(newTeams) => (selected = newTeams)}
		/>
	</form>
</div>
