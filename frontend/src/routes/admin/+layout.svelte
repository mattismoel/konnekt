<script lang="ts">
	import '../../app.css';

	import { cn } from '$lib/clsx';
	import { toaster } from '$lib/toaster.svelte';

	import Toast from '$lib/components/Toast.svelte';
	import Sidebar from './Sidebar.svelte';
	import { MediaQuery } from 'svelte/reactivity';

	let { children, data } = $props();
	let { user, roles } = $derived(data);

	let windowWidth = $state(0);

	const large = new MediaQuery('min-width: 768px');

	let sidebarExpanded = $derived(large.current);
</script>

<svelte:window bind:innerWidth={windowWidth} />

<main class="bg-background flex min-h-svh w-screen">
	<Sidebar
		{user}
		{roles}
		expanded={sidebarExpanded}
		onToggle={() => (sidebarExpanded = !sidebarExpanded)}
	/>
	<div
		class={cn('grid-cols grid flex-1', {
			'pl-sidenav-lg': sidebarExpanded,
			'pl-sidenav-sm': !sidebarExpanded
		})}
	>
		{@render children()}
	</div>
	<div class="fixed right-4 bottom-4 z-50 flex flex-col gap-2">
		{#each toaster.toasts as toast (toast.id)}
			<Toast {toast} onDelete={() => toaster.removeToast(toast.id)} />
		{/each}
	</div>
</main>
