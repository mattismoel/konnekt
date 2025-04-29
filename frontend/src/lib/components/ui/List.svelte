<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLUListElement>;

	let { children, ...rest }: Props = $props();
</script>

<ul {...rest} class="space-y-2">
	{@render children?.()}
</ul>
