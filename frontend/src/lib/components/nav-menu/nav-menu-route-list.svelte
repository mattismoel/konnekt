<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	let { children, ...rest }: HTMLAttributes<HTMLUListElement> = $props();
</script>

<ul {...rest} class="flex flex-col gap-4">
	{@render children?.()}
</ul>
