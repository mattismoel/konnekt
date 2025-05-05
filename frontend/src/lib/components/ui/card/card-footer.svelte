<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement>;

	let { children, ...rest }: Props = $props();
</script>

<div {...rest} class="p-6">{@render children?.()}</div>
