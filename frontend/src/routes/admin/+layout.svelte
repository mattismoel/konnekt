<script lang="ts">
	import '../../app.css';

	import { toaster } from '$lib/toaster.svelte';

	import Toast from '$lib/components/Toast.svelte';
	import Sidebar from './Sidebar.svelte';
	import * as Navbar from '$lib/components/navbar/index';
	import { beforeNavigate } from '$app/navigation';
	import Logo from '$lib/assets/Logo.svelte';
	import MenuIcon from '~icons/mdi/menu';
	import { authStore } from '$lib/auth.svelte';

	let { children, data } = $props();

	let sidebarExpanded = $state(false);

	beforeNavigate(() => {
		sidebarExpanded = false;
	});
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

			<a href="/admin/members/{data.member.id}" class="group">
				<img
					src={data.member.profilePictureUrl}
					alt="Profile"
					class="h-8 w-8 rounded-full object-cover outline outline-zinc-700 group-hover:outline-2"
				/>
			</a>
		</Navbar.Content>
	</Navbar.Root>
	<Sidebar bind:expanded={sidebarExpanded} />

	{@render children()}

	<div class="fixed right-4 bottom-4 z-50 flex flex-col gap-2">
		{#each toaster.toasts as toast (toast.id)}
			<Toast {toast} onDelete={() => toaster.removeToast(toast.id)} />
		{/each}
	</div>
</div>
