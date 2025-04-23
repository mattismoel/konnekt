<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Avatar from '$lib/assets/avatar.png';
	import ListEntry from '$lib/components/ListEntry.svelte';
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

	// TODO: Implement...
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

<ListEntry class="gap-4">
	<div class="flex flex-1 items-center gap-4">
		<img
			src={member.profilePictureUrl || Avatar}
			alt="Profil"
			class="h-8 w-8 rounded-full object-cover"
		/>
		<span class="line-clamp-1">{member.firstName} {member.lastName} </span>
	</div>
	<MemberStatusIndicator status="non-approved" class="hidden md:block" />
	<div class="text-text/75">
		<button class="p-1 hover:text-green-500" onclick={approve}><ApproveIcon /></button>
		<button class="p-1 hover:text-red-500" onclick={disapprove}><DisapproveIcon /></button>
	</div>
</ListEntry>
