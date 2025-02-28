<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDialogElement> & {
		show: boolean;
	};

	let { children, show }: Props = $props();

	let dialog: HTMLDialogElement;

	$effect(() => (show ? dialog.showModal() : dialog.close()));
</script>

<dialog
	bind:this={dialog}
	class="fixed top-1/2 left-1/2 min-w-96 -translate-x-1/2 -translate-y-1/2 flex-col
		overflow-hidden rounded-md border border-zinc-800"
>
	{@render children?.()}
</dialog>
