<script lang="ts">
	import { goto } from '$app/navigation';
	import type { Artist } from '$lib/features/artist/artist';
	import { hasPermissions, type Permission } from '$lib/features/auth/permission';
	import ContextMenu from '$lib/components/ui/context-menu/ContextMenu.svelte';
	import ContextMenuEntry from '$lib/components/ui/context-menu/ContextMenuEntry.svelte';
	import MenuIcon from '~icons/mdi/dots-vertical';
	import ListEntry from '$lib/components/ui/ListEntry.svelte';

	type Props = {
		artist: Artist;
		onDelete: () => void;
		memberPermissions: Permission[];
	};

	let { artist, memberPermissions, onDelete }: Props = $props();

	let showContextMenu = $state(false);
</script>

<ListEntry>
	<a href="/admin/artists/edit/{artist.id}" class="flex w-full flex-col">
		<span>{artist.name}</span>
		<span class="text-text/50">{artist.genres.map((genre) => genre.name).join(', ')}</span>
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
			disabled={!hasPermissions(memberPermissions, ['edit:artist'])}
			href="admin/artists/edit/{artist.id}"
		>
			Redigér
		</ContextMenuEntry>
		<ContextMenuEntry
			disabled={!hasPermissions(memberPermissions, ['delete:artist'])}
			action={onDelete}>Slet</ContextMenuEntry
		>
	</ContextMenu>
</ListEntry>
