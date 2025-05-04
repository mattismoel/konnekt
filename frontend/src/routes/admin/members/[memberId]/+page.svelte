<script lang="ts">
	import { hasPermissions } from '$lib/features/auth/permission';

	import MemberStatusIndicator from '$lib/components/MemberStatusIndicator.svelte';
	import MemberForm from './MemberForm.svelte';
	import MemberTeamsForm from './MemberTeamsForm.svelte';
	import ProfilePictureSelector from '$lib/components/ProfilePictureSelector.svelte';

	let { data } = $props();
</script>

<main class="px-auto space-y-16 py-16 pt-32">
	<MemberForm member={data.member} teams={data.teams} />
	{#if hasPermissions(data.currentMember.permissions, ['edit:member'])}
		<MemberTeamsForm member={data.member} teams={data.teams} defaultSelected={data.member.teams} />
	{/if}
</main>
