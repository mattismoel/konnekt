<script lang="ts">
	import type { WithChildren } from '$lib/types';

	type Props = WithChildren<{}>;

	let { children }: Props = $props();
</script>

<div class="flex items-center gap-8">
	{@render children?.()}
</div>
