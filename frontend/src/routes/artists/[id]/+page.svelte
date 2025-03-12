<script lang="ts">
	import Fader from '$lib/components/ui/Fader.svelte';
	import { socialUrlToIcon } from '$lib/social';
	import { trackIdFromUrl } from '$lib/spotify.js';

	let { data } = $props();
	let { artist } = $derived(data);

	let trackId = $derived(trackIdFromUrl(artist.previewUrl));
</script>

<main>
	<div class="grid min-h-svh grid-cols-2">
		<div class="relative isolate h-full">
			<h1 class="absolute top-20 left-24 z-10 w-screen text-9xl font-bold">{artist.name}</h1>
			<img src={artist.imageUrl} alt="Cover af {artist.name}" class="h-full object-cover" />
			<Fader direction="left" class="absolute w-64 from-zinc-950" />
		</div>
		<article class="prose prose-invert flex flex-col gap-8 pt-64 pb-16 pl-24">
			<div class="flex-1">
				{@html artist.description}
			</div>
			{#if trackId}
				{@render spotifyPreview(trackId)}
			{/if}
			{@render socials(artist.socials)}
		</article>
	</div>
</main>

{#snippet spotifyPreview(trackId: string)}
	{@const src = `https://open.spotify.com/embed/track/${trackId}?utm_source=generator&theme=0`}
	<iframe
		title="Audio preview"
		style:border-radius="12px"
		{src}
		width="100%"
		height="152"
		frameBorder="0"
		allowfullscreen
		allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
		loading="lazy"
	></iframe>
{/snippet}

{#snippet socials(urls: string[])}
	<div class="flex w-full justify-end gap-4 text-2xl">
		{#each urls as url (url)}
			{@const Icon = socialUrlToIcon(url)}
			<a href={url}>
				<Icon />
			</a>
		{/each}
	</div>
{/snippet}
