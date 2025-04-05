<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLLIElement> & {
		action: () => void;
	};

	let { action, children }: Props = $props();
</script>

<li class="text-text/75 hover:text-text py-2">
	<button onclick={action}>
		{@render children?.()}
	</button>
</li>
