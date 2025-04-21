<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import Avatar from '$lib/assets/avatar.png';
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

<li class="flex justify-between rounded-sm p-3 hover:bg-zinc-900">
	<div class="flex items-center gap-4">
		<img src={member.profilePictureUrl || Avatar} alt="Profil" class="h-8 w-8 rounded-full" />
		<span>{member.firstName} {member.lastName} </span>
	</div>
	<div class="text-text/75">
		<button class="p-1 hover:text-green-500" onclick={approve}><ApproveIcon /></button>
		<button class="p-1 hover:text-red-500" onclick={disapprove}><DisapproveIcon /></button>
	</div>
</li>
