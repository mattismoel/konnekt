<script lang="ts">
	import type { WithChildren } from '$lib/types';

	type Props = WithChildren<{}>;

	let { children }: Props = $props();
</script>

<div class="flex items-center gap-6">
	{@render children?.()}
</div>
