<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		srcs: Map<string, string>;
	};

	let { srcs }: Props = $props();
</script>

<div class="group flex gap-8 overflow-hidden">
	{#each { length: 2 } as _, i}
		{@render track(i)}
	{/each}
</div>

{#snippet track(idx: number)}
	<div class="animate-slide flex h-8 space-x-8" aria-hidden={idx > 0}>
		{#each srcs as [name, src] (src)}
			<img {src} alt={name} class="max-w-none" />
		{/each}
	</div>
{/snippet}
