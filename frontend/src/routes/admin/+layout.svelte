<script lang="ts">
	import Toast from '$lib/components/ui/Toast.svelte';
	import { toaster } from '$lib/toaster.svelte';

	let { children } = $props();
</script>

{@render children?.()}
<div class="fixed right-4 bottom-4 z-50 flex flex-col gap-2">
	{#each toaster.toasts as toast (toast.id)}
		<Toast {toast} onDelete={() => toaster.removeToast(toast.id)} />
	{/each}
</div>
