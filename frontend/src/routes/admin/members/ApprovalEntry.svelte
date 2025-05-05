<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Avatar from '$lib/assets/avatar.png';
	import * as List from '$lib/components/ui/list/index';
	import MemberStatusIndicator from '$lib/components/MemberStatusIndicator.svelte';
	import { approveMember, deleteMember, type Member } from '$lib/features/auth/member';
	import { toaster } from '$lib/toaster.svelte';
	import ApproveIcon from '~icons/mdi/account-check';
	import DisapproveIcon from '~icons/mdi/trash';

	type Props = {
		member: Member;
	};

	let { member }: Props = $props();

	const approve = async () => {
		try {
			await approveMember(fetch, member.id);
			toaster.addToast('Bruger godkendt');
			await invalidateAll();
		} catch (e) {
			toaster.addToast('Kunne ikke godkende bruger', 'Noget gik galt...', 'error');
		}
	};

	const disapprove = async () => {
		try {
			await deleteMember(fetch, member.id);
			toaster.addToast('Bruger forkastet');
			await invalidateAll();
		} catch (e) {
			toaster.addToast('Kunne ikke forkaste bruger', 'Noget gik galt...', 'error');
		}
	};
</script>

<List.Entry class="gap-4">
	<List.Section href="/admin/members/{member.id}">
		<div class="flex flex-1 items-center gap-4">
			<img
				src={member.profilePictureUrl || Avatar}
				alt="Profil"
				class="h-8 w-8 rounded-full object-cover"
			/>
			<span class="line-clamp-1">{member.firstName} {member.lastName} </span>
		</div>
	</List.Section>

	<List.Section expand={false} class="flex-row gap-6">
		<MemberStatusIndicator status="non-approved" class="hidden md:block" />
		<div class="text-text/75 flex gap-2">
			<button class="p-1 hover:text-green-500" onclick={approve}><ApproveIcon /></button>
			<button class="p-1 hover:text-red-500" onclick={disapprove}><DisapproveIcon /></button>
		</div>
	</List.Section>
</List.Entry>
