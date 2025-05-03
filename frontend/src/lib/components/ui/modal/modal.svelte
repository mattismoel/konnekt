<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDialogElement> & {
		show: boolean;
	};

	let { children, show = $bindable(), ...rest }: Props = $props();

	let dialog: HTMLDialogElement;

	$effect(() => {
		show ? dialog.showModal() : dialog.close();
	});
</script>

<dialog
	onclose={() => (show = false)}
	onclick={(e) => e.target === dialog && dialog.close()}
	bind:this={dialog}
	class={cn(
		'fixed top-1/2 left-1/2 w-full min-w-xs -translate-x-1/2 -translate-y-1/2 flex-col overflow-hidden rounded-md border border-zinc-800 sm:min-w-lg',
		rest.class
	)}
>
	{@render children?.()}
</dialog>
