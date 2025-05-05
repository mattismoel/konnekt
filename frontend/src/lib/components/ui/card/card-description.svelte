<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLParagraphElement>;

	let { children, ...rest }: Props = $props();
</script>

<p {...rest} class="text-text/50">{@render children?.()}</p>
