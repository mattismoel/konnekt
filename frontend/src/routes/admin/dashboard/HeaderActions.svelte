<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement>;

	let { children }: Props = $props();
</script>

<div class="flex gap-4">
	{@render children?.()}
</div>
