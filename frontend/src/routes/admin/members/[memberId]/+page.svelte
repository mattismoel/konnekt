<script lang="ts">
	import { authStore } from '$lib/auth.svelte';
	import { hasPermissions } from '$lib/features/auth/permission';

	import MemberForm from './MemberForm.svelte';
	import MemberTeamsForm from './MemberTeamsForm.svelte';

	let { data } = $props();
</script>

<main class="flex flex-col gap-16">
	<MemberForm member={data.member} teams={data.teams} />
	{#if hasPermissions(authStore.permissions, ['edit:member'])}
		<MemberTeamsForm member={data.member} teams={data.allTeams} defaultSelected={data.teams} />
	{/if}
</main>
