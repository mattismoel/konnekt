<script lang="ts">
	import { hasPermissions } from '$lib/features/auth/permission';

	import MemberForm from './MemberForm.svelte';
	import MemberTeamsForm from './MemberTeamsForm.svelte';

	let { data } = $props();
</script>

<main class="flex flex-col gap-16">
	<MemberForm member={data.member} teams={data.teams} />
	{#if hasPermissions(data.currentMember.permissions, ['edit:member'])}
		<MemberTeamsForm member={data.member} teams={data.teams} defaultSelected={data.member.teams} />
	{/if}
</main>
