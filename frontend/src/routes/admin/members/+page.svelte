<script lang="ts">
	import DashboardHeader from '../DashboardHeader.svelte';
	import DashboardLayout from '../DashboardLayout.svelte';
	import MemberList from './MemberList.svelte';

	let { data } = $props();
</script>

<DashboardLayout>
	<DashboardHeader
		title="Medlemmer"
		description="Her kan du administrere alle medlemmer af foreningen, herunder deres roller i foreningen, og
			dermed deres tilladelser."
	></DashboardHeader>

	<MemberList members={data.members} pending={data.pending} />
</DashboardLayout>
