<script lang="ts">
	import List from '$lib/components/ui/List.svelte';
	import type { Member } from '$lib/features/auth/member';
	import type { Permission } from '$lib/features/auth/permission';
	import ApprovalEntry from './ApprovalEntry.svelte';
	import MemberEntry from './MemberEntry.svelte';

	type Props = {
		members: Member[];
		pending: Member[];

		memberPermissions: Permission[];
	};

	let { members, pending, memberPermissions }: Props = $props();
</script>

<div class="space-y-8">
	{#if pending.length > 0}
		<section>
			<h1 class="mb-4">Anmodninger ({pending.length})</h1>
			<List>
				{#each pending as member (member.id)}
					<ApprovalEntry {member} />
				{/each}
			</List>
		</section>
	{/if}

	<section>
		<h1 class="mb-4">Medlemmer</h1>
		<List class="space-y-2">
			{#each members as member (member.id)}
				<MemberEntry {member} {memberPermissions} />
			{/each}
		</List>
	</section>
</div>
