<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLLIElement> & {
		action: () => void;
		disabled: boolean;
	};

	let { action, disabled, children }: Props = $props();
</script>

<li class="text-text/85 hover:text-text py-2">
	<button {disabled} onclick={action} class="disabled:text-text/50">
		{@render children?.()}
	</button>
</li>
