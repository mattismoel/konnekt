<script lang="ts">
	import type { Member, Team, TeamType } from "$lib/features/auth/auth";
	import { getTeams } from "$lib/features/auth/auth.remote";

	const includedTeamNames: TeamType[] = [
		"project-leader",
		"public-relations",
		"event-management",
		"booking",
		"visual-identity",
		"economy"
	];

	type Props = {
		members: Member[];
	};

	let { members }: Props = $props();

	const { items: teams } = $derived(await getTeams(undefined));

	const includedTeams = $derived(teams.filter((team) => includedTeamNames.includes(team.name)));

	const includedMembers = $derived(
		members.filter((member) =>
			member.teams.some((team) =>
				includedTeams.some((includedTeam) => includedTeam.name === team.name)
			)
		)
	);

	const specialMembers = $derived(includedMembers.filter((member) => member.specialRole));
	const regularMembers = $derived(includedMembers.filter((member) => !member.specialRole));

	const extractValidTeamNames = (teams: Team[]) => {
		return teams
			.filter((team) => includedTeams.some((t) => t.name === team.name))
			.map((t) => t.name);
	};
</script>

<div class="@container flex flex-col">
	<div class="mb-16 grid grid-cols-1 gap-2 @3xl:grid-cols-2">
		{#each specialMembers as { firstName, lastName, avatar, specialRole, teams }}
			<div
				class="group flex items-center overflow-hidden rounded-full border border-zinc-800 bg-zinc-900 p-1 transition-colors"
			>
				<img
					src={avatar}
					alt="{firstName} {lastName}"
					class="aspect-square w-24 shrink-0 rounded-full object-cover @lg:w-32"
				/>
				<div class="h-mi flex w-full flex-col gap-2 overflow-hidden px-8 py-4">
					<span class="line-clamp-1 font-medium text-text-light @lg:text-xl">
						{firstName}
						{lastName}
					</span>
					<div class="w-full">
						<span class="line-clamp-1">
							{#if specialRole}
								{specialRole}
							{:else}
								{extractValidTeamNames(teams).join(", ")}
							{/if}
						</span>
					</div>
				</div>
			</div>
		{/each}
	</div>

	{#if regularMembers.length > 0}
		<h4 class="text-text/50 mb-12 text-center">
			Øvrige {#if regularMembers.length > 5}{regularMembers.length}{/if} medlemmer
		</h4>
	{/if}

	<div class="flex flex-wrap justify-center gap-x-20 gap-y-12">
		{#each regularMembers as member}
			{@const memberTeams = member.teams.filter((team) =>
				includedTeams.some((t) => t.name === team.name)
			)}
			{@const teamsString = memberTeams.map((t) => t.displayName).join(", ")}

			<div class="flex flex-col items-center">
				<p class="mb-1 text-base font-medium">
					{member.firstName}
					{member.lastName}
				</p>
				<p class="text-text/50 text-xs">{teamsString}</p>
			</div>
		{/each}
	</div>
</div>
