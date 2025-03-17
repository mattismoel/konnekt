<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';
	import { cn } from '$lib/clsx';

	type Props = HTMLAttributes<HTMLDivElement> & {
		srcs: Map<string, string>;
	};

	let { srcs, ...rest }: Props = $props();
</script>

<div class={cn('relative isolate h-8 w-full overflow-hidden whitespace-nowrap', rest.class)}>
	{@render track(srcs)}
	{@render track(srcs)}
	{@render track(srcs)}
</div>

{#snippet track(srcs: Map<string, string>)}
	<div class="animate-slide inline-block h-full w-fit flex-nowrap">
		{#each [...srcs] as [name, src] (src)}
			<img {src} alt={name} class="mx-8 inline h-full" />
		{/each}
	</div>
{/snippet}
