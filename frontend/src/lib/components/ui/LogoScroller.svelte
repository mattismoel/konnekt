<script lang="ts">
	import Fader from './Fader.svelte';

	type Props = {
		srcs: Map<string, string>;
	};

	let { srcs }: Props = $props();
</script>

<div class="relative isolate w-full overflow-hidden whitespace-nowrap">
	<Fader direction="right" class="absolute z-50 w-32 from-zinc-900" />
	<Fader direction="left" class="absolute z-50 w-32 from-zinc-900" />
	{@render track(srcs)}
	{@render track(srcs)}
</div>

{#snippet track(srcs: Map<string, string>)}
	<div class="animate-slide inline-block w-fit flex-nowrap">
		{#each [...srcs] as [name, src] (src)}
			<img {src} alt={name} class="mx-8 inline h-16" />
		{/each}
	</div>
{/snippet}
