<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Artist } from '$lib/features/artist/artist';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import MenuIcon from '~icons/mdi/dots-vertical';

	type Props = {
		artist: Artist;
		onDelete: () => void;
		userPermissions: Permission[];
	};

	let { artist, userPermissions, onDelete }: Props = $props();

	let showContextMenu = $state(false);
</script>

<li class="relative flex">
	<a
		href="/admin/artists/edit?id={artist.id}"
		class="flex flex-1 items-center rounded-sm border border-transparent px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
	>
		<span class="flex-1 font-medium">{artist.name}</span>
	</a>
	<button onclick={() => (showContextMenu = true)} class="rounded-md p-2 hover:bg-zinc-900">
		<MenuIcon />
	</button>
	<ContextMenu
		open={showContextMenu}
		onClose={() => (showContextMenu = false)}
		class="absolute top-1/2 right-4"
	>
		<ContextMenuEntry
			disabled={!hasPermissions(userPermissions, ['delete:artist'])}
			action={onDelete}>Slet</ContextMenuEntry
		>
		<ContextMenuEntry
			disabled={!hasPermissions(userPermissions, ['edit:artist'])}
			action={() => goto(`/admin/artists/edit?id=${artist.id}`)}>Redigér</ContextMenuEntry
		>
	</ContextMenu>
</li>
