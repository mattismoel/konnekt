<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLParagraphElement>;
	let { children, ...rest }: Props = $props();
</script>

<p class="text-text/50 text-sm" {...rest}>{@render children?.()}</p>
