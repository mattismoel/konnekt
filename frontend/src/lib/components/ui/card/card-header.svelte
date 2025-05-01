<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement>;

	let { children, ...rest }: Props = $props();
</script>

<div {...rest} class="flex flex-col gap-2 p-6">
	{@render children?.()}
</div>
