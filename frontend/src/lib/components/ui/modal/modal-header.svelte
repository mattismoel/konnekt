<script lang="ts">
	import type { WithChildren } from '$lib/types';
	import CloseIcon from '~icons/mdi/close';

	type Props = WithChildren<{
		onClose: () => void;
	}>;

	let { children, onClose }: Props = $props();
</script>

<div class="text-text relative flex flex-col justify-center gap-2 bg-zinc-950 p-6">
	{@render children?.()}
	<button
		type="button"
		onclick={onClose}
		class="text-text/50 hover:text-text absolute top-6 right-6"><CloseIcon /></button
	>
</div>
