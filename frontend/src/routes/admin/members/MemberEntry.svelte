<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import * as ContextMenu from '$lib/components/ui/context-menu/index';
	import AvatarImage from '$lib/assets/avatar.png';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';
	import MemberStatusIndicator from '$lib/components/MemberStatusIndicator.svelte';
	import { deleteMember, type Member } from '$lib/features/auth/member';
	import { hasPermissions } from '$lib/features/auth/permission';
	import { toaster } from '$lib/toaster.svelte';
	import { authStore } from '$lib/auth.svelte';

	type Props = {
		member: Member;
	};

	let { member }: Props = $props();

	let isContextMenuOpen = $state(false);

	let fullName = $derived(`${member.firstName} ${member.lastName}`);

	const handleDeleteMember = async () => {
		if (
			!confirm(
				`Er du sikker på at du vil slette ${fullName} fra foreningen?\n\nHandlingen kan ikke fortrydes.`
			)
		) {
			return;
		}
		try {
			await deleteMember(fetch, member.id);
			toaster.addToast('Medlem slettet');
			await invalidateAll();
		} catch (e) {
			toaster.addToast('Kunne ikke slette medlemmet', 'Noget gik galt...', 'error');
		}
	};
</script>

<ListEntry class="relative">
	<div class="flex flex-1 items-center gap-4">
		<img
			src={member.profilePictureUrl || AvatarImage}
			alt="Profil"
			class="h-8 w-8 rounded-full object-cover"
		/>
		<span class="line-clamp-1">{member.firstName} {member.lastName}</span>
	</div>
	<MemberStatusIndicator
		status={member.active ? 'approved' : 'non-approved'}
		class="hidden md:flex"
	/>
	<ContextMenu.Button onclick={() => (isContextMenuOpen = true)} />
	<ContextMenu.Root bind:show={isContextMenuOpen}>
		<ContextMenu.Entry
			href="/admin/members/{member.id}"
			disabled={!hasPermissions(authStore.permissions, ['edit:member'])}
		>
			Redigér
		</ContextMenu.Entry>
		<ContextMenu.Entry
			onclick={handleDeleteMember}
			disabled={!hasPermissions(authStore.permissions, ['delete:member'])}
		>
			Slet
		</ContextMenu.Entry>
	</ContextMenu.Root>
</ListEntry>
