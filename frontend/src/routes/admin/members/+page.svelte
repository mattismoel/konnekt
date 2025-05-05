<script lang="ts">
	import * as AdminHeader from '$lib/components/admin-header/index';
	import MemberList from './MemberList.svelte';

	let { data } = $props();
</script>

<AdminHeader.Root>
	<AdminHeader.Title>Medlemmer</AdminHeader.Title>
	<AdminHeader.Description
		>Her kan du administrere alle medlemmer af foreningen, herunder deres roller i foreningen, og
		dermed deres tilladelser."</AdminHeader.Description
	>
</AdminHeader.Root>

<main class="pt-16">
	<MemberList members={data.members} pending={data.pending} />
</main>
