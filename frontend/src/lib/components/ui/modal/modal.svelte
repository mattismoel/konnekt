<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDialogElement> & {
		show: boolean;
	};

	let { children, show = $bindable(), ...rest }: Props = $props();

	let dialog: HTMLDialogElement;

	$effect(() => {
		if (show) dialog.showModal();
	});
</script>

<dialog
	onclose={() => (show = false)}
	onclick={(e) => e.target === dialog && dialog.close()}
	bind:this={dialog}
	class={cn(
		'fixed top-1/2 left-1/2 min-w-96 -translate-x-1/2 -translate-y-1/2 flex-col overflow-hidden rounded-md border border-zinc-800',
		rest.class
	)}
>
	{@render children?.()}
</dialog>
