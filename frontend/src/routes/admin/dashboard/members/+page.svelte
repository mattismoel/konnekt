<script lang="ts">
	import MemberList from './MemberList.svelte';

	let { data } = $props();
</script>

<main class="space-y-16 px-8 py-16 md:px-16">
	<div>
		<h1 class="mb-4 text-4xl font-bold">Medlemmer</h1>
		<span class="text-text/50"
			>Her kan du administrere alle medlemmer af foreningen, herunder deres roller i foreningen, og
			dermed deres tilladelser.</span
		>
	</div>
	<MemberList
		members={data.members}
		pending={data.pending}
		memberPermissions={data.member.permissions}
	/>
</main>
