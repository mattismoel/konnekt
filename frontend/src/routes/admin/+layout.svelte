<script lang="ts">
	import '../../app.css';

	import { toaster } from '$lib/toaster.svelte';

	import Toast from '$lib/components/Toast.svelte';
	import Sidebar from './Sidebar.svelte';
	import * as Navbar from '$lib/components/navbar/index';
	import { beforeNavigate, goto } from '$app/navigation';
	import Logo from '$lib/assets/Logo.svelte';
	import MenuIcon from '~icons/mdi/menu';
	import { authStore } from '$lib/auth.svelte';
	import * as ContextMenu from '$lib/components/ui/context-menu/index';
	import { logOut } from '$lib/features/auth/auth';
	import { APIError } from '$lib/api';

	let { children } = $props();

	let sidebarExpanded = $state(false);
	let userContextMenuOpen = $state(false);

	beforeNavigate(() => {
		sidebarExpanded = false;
	});

	const handleLogOut = async () => {
		try {
			await logOut(fetch);
			await goto('/');
		} catch (e) {
			if (e instanceof APIError) {
				toaster.addToast('Kunne ikke logge ud...', e.cause, 'error');
			}
		}
	};
</script>

<div class="bg-background flex min-h-svh w-screen flex-col">
	<Navbar.Root>
		<Navbar.Header>
			<button
				onclick={() => (sidebarExpanded = true)}
				class="text-text/75 hover:text-text text-xl md:hidden"><MenuIcon /></button
			>
			<a href="/">
				<Logo class="h-4" />
			</a>
		</Navbar.Header>

		<Navbar.Content>
			<Navbar.RouteList>
				<Navbar.RouteEntry pathname="/admin/events" name="Events" />
				<Navbar.RouteEntry pathname="/admin/artists" name="Kunstnere" />
				<Navbar.RouteEntry pathname="/admin/venues" name="Venues" />
				<Navbar.RouteEntry pathname="/admin/members" name="Medlemmer" />
			</Navbar.RouteList>

			<button class="group" onclick={() => (userContextMenuOpen = !userContextMenuOpen)}>
				<img
					src={authStore.member?.profilePictureUrl}
					alt="Profile"
					class="h-8 w-8 rounded-full object-cover outline outline-zinc-700 group-hover:outline-2"
				/>
				<ContextMenu.Root bind:show={userContextMenuOpen}>
					<ContextMenu.Entry disabled={false} href="/admin/members/{authStore.member?.id}">
						Redigér
					</ContextMenu.Entry>
					<ContextMenu.Entry disabled={false} onclick={handleLogOut}>Log ud</ContextMenu.Entry>
				</ContextMenu.Root>
			</button>
		</Navbar.Content>
	</Navbar.Root>

	<Sidebar bind:expanded={sidebarExpanded} />

	<div class="px-auto py-16 pt-32">
		{@render children()}
	</div>

	<div class="fixed right-4 bottom-4 z-50 flex flex-col gap-2">
		{#each toaster.toasts as toast (toast.id)}
			<Toast {toast} onDelete={() => toaster.removeToast(toast.id)} />
		{/each}
	</div>
</div>
