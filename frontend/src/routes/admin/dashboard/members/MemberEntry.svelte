<script lang="ts">
	import type { Member } from '$lib/features/auth/member';

	type Props = {
		member: Member;
	};

	let { member }: Props = $props();
</script>

<li class="rounded-sm bg-zinc-900 p-3">
	<img src={member.profilePictureUrl} alt="" />
	{member.firstName}
	{member.lastName}
</li>
