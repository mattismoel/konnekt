<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLHeadingElement>;

	let { children, ...rest }: Props = $props();
</script>

<h1 class="font-medium" {...rest}>{@render children?.()}</h1>
