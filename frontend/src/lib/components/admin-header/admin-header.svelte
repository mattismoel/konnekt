<script lang="ts">
	import type { WithChildren } from '$lib/types';

	type Props = WithChildren<{}>;

	let { children }: Props = $props();
</script>

<header class="flex flex-col gap-6">
	{@render children?.()}
</header>
