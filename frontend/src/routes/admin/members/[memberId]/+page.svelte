<script lang="ts">
	import { hasPermissions } from '$lib/features/auth/permission';
	import AvatarImage from '$lib/assets/avatar.png';

	import MemberStatusIndicator from '$lib/components/MemberStatusIndicator.svelte';
	import MemberForm from './MemberForm.svelte';
	import MemberTeamsForm from './MemberTeamsForm.svelte';
	import ProfilePictureSelector from '$lib/components/ProfilePictureSelector.svelte';

	let { data } = $props();
</script>

<main class="space-y-16 px-4 py-16 sm:px-16">
	{@render header()}
	<MemberForm member={data.member} currentMember={data.currentMember} teams={data.teams} />
	{#if hasPermissions(data.currentMember.permissions, ['edit:member'])}
		<MemberTeamsForm member={data.member} teams={data.teams} defaultSelected={data.member.teams} />
	{/if}
</main>

{#snippet header()}
	{@const fullName = `${data.member.firstName} ${data.member.lastName}`}

	<header class="flex flex-col items-center gap-8 md:flex-row">
		<form action="" class="">
			<ProfilePictureSelector file={null} imageUrl={data.member.profilePictureUrl} />
		</form>
		<div class="flex flex-col items-center space-y-4 md:items-start">
			<div class="flex flex-col items-center space-y-1 md:items-start">
				<h1 class="text-2xl font-semibold">{fullName}</h1>
				<span class="text-text/50 text-center md:text-left"
					>{data.member.teams.map(({ displayName }) => displayName).join(', ')}</span
				>
			</div>
			<MemberStatusIndicator status={data.member.active ? 'approved' : 'non-approved'} />
		</div>
	</header>
{/snippet}
