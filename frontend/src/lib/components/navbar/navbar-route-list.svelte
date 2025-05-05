<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLUListElement>;

	let { children, ...rest }: Props = $props();
</script>

<ul class={cn('hidden items-center gap-6 text-lg text-zinc-50 md:flex', rest.class)}>
	{@render children?.()}
</ul>
