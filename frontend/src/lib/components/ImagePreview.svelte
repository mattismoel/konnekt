<script lang="ts">
	type Props = {
		src: string | undefined;
	};

	let { src = $bindable() }: Props = $props();
</script>

<div
	class="flex aspect-video items-center justify-center overflow-hidden rounded-xl border border-zinc-800"
>
	{#if src}
		<img {src} alt="Preview" class="h-full w-full object-cover" />
	{:else}
		<p>Vælg et billede...</p>
	{/if}
</div>
