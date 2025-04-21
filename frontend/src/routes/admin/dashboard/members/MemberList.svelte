<script lang="ts">
	import type { Member } from '$lib/features/auth/member';
	import ApprovalEntry from './ApprovalEntry.svelte';
	import MemberEntry from './MemberEntry.svelte';

	type Props = {
		members: Member[];
		pending: Member[];
	};

	let { members, pending }: Props = $props();
</script>

<div class="space-y-8">
	{#if pending.length > 0}
		<section>
			<h1 class="mb-4">Anmodninger ({pending.length})</h1>
			<ul class="space-y-2">
				{#each pending as member (member.id)}
					<ApprovalEntry {member} />
				{/each}
			</ul>
		</section>
	{/if}

	<section>
		<h1 class="mb-4">Medlemmer</h1>
		<ul class="space-y-2">
			{#each members as member (member.id)}
				<MemberEntry {member} />
			{/each}
		</ul>
	</section>
</div>
