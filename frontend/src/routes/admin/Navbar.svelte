<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import MenuIcon from '~icons/mdi/menu';

	type Props = {
		onSidebarOpen: () => void;
	};

	let { onSidebarOpen }: Props = $props();
</script>

<nav class="h-nav px-auto py-2">
	<Button variant="secondary" type="button" onclick={onSidebarOpen}><MenuIcon /></Button>
</nav>
