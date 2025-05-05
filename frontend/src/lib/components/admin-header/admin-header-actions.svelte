<script lang="ts">
	import type { WithChildren } from '$lib/types';

	type Props = WithChildren<{}>;

	let { children }: Props = $props();
</script>

<div class="flex gap-4">
	{@render children?.()}
</div>
