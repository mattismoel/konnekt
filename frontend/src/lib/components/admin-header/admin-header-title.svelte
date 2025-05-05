<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLHeadingElement>;

	let { children, ...rest }: Props = $props();
</script>

<h1 {...rest} class="font-heading text-4xl font-bold">{@render children?.()}</h1>
