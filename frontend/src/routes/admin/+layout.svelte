<script lang="ts">
	import '../../app.css';

	import { toaster } from '$lib/toaster.svelte';

	import Toast from '$lib/components/Toast.svelte';
	import Sidebar from './Sidebar.svelte';
	import { MediaQuery } from 'svelte/reactivity';
	import Navbar from './Navbar.svelte';

	let { children, data } = $props();
	let { member } = $derived(data);

	let windowWidth = $state(0);

	const large = new MediaQuery('min-width: 768px');

	let sidebarExpanded = $derived(large.current);
</script>

<svelte:window bind:innerWidth={windowWidth} />

<main class="bg-background flex min-h-svh w-screen flex-col">
	<Navbar onSidebarOpen={() => (sidebarExpanded = true)} />
	<Sidebar
		{member}
		expanded={sidebarExpanded}
		onToggle={() => (sidebarExpanded = !sidebarExpanded)}
	/>
	<div class="grid flex-1 grid-cols-1 overflow-hidden">
		{@render children()}
	</div>
	<div class="fixed right-4 bottom-4 z-50 flex flex-col gap-2">
		{#each toaster.toasts as toast (toast.id)}
			<Toast {toast} onDelete={() => toaster.removeToast(toast.id)} />
		{/each}
	</div>
</main>
