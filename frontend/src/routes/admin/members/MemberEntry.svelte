<script lang="ts">
	import { invalidateAll } from '$app/navigation';
	import * as ContextMenu from '$lib/components/ui/context-menu/index';
	import AvatarImage from '$lib/assets/avatar.png';
	import * as List from '$lib/components/ui/list/index';
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

<List.Entry class="relative">
	<List.Section href="/admin/members/{member.id}" class="flex-row items-center gap-4">
		<img
			src={member.profilePictureUrl || AvatarImage}
			alt="Profil"
			class="h-8 w-8 rounded-full object-cover"
		/>
		<span class="line-clamp-1">{member.firstName} {member.lastName}</span>
	</List.Section>

	<List.Section class="flex-row items-center gap-4" expand={false}>
		<MemberStatusIndicator
			status={member.active ? 'approved' : 'non-approved'}
			class="hidden md:flex"
		/>
		<ContextMenu.Button onclick={() => (isContextMenuOpen = true)} />
	</List.Section>

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
</List.Entry>
