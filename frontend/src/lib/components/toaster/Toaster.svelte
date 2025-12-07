<script lang="ts">
	import { fade } from "svelte/transition";
	import { getToastContext, type Toast } from "./toaster.svelte.ts";

	const toaster = getToastContext();
</script>

<div class="fixed bottom-8 flex min-h-4 w-full justify-end">
	<ul class="mx-responsive flex w-full flex-col items-end gap-1">
		{#each toaster.toasts as toast}
			{@render toastEntry(toast)}
		{/each}
	</ul>
</div>

{#snippet toastEntry(toast: Toast)}
	<li
		transition:fade={{ duration: 200 }}
		class={[
			"flex min-w-sm items-center justify-between gap-8 rounded-2xl border px-6 py-2 font-medium",
			toast.severity === "info" && "border-blue-900 bg-blue-950 text-blue-300",
			toast.severity === "warning" && "border-amber-900 bg-amber-950 text-amber-500",
			toast.severity === "dangerous" && "border-red-900 bg-red-950 text-red-500"
		]}
	>
		<div class="flex items-center gap-4">
			<p>
				{#if toast.severity === "info"}
					I
				{:else if toast.severity === "warning"}
					W
				{:else if toast.severity === "dangerous"}
					D
				{/if}
			</p>
			{toast.text}
		</div>
		<button type="button" onclick={() => toaster.remove(toast.id)}>X</button>
	</li>
{/snippet}
