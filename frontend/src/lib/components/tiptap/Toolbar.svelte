<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement>;

	let { children, ...rest }: Props = $props();
</script>

<div class="flex gap-4 overflow-hidden rounded-t-md border border-zinc-800 bg-zinc-900 p-1">
	{@render children?.()}
</div>
