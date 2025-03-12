<script lang="ts">
	import { cn } from '$lib/clsx';
	import EventCard from '$lib/components/EventCard.svelte';
	import EventCaroussel from '$lib/components/EventCaroussel.svelte';
	import SpotifyPreview from '$lib/components/SpotifyPreview.svelte';
	import Fader from '$lib/components/ui/Fader.svelte';
	import { socialUrlToIcon } from '$lib/social';
	import { trackIdFromUrl } from '$lib/spotify.js';

	let { data } = $props();
	let { artist, events } = $derived(data);

	let trackId = $derived(trackIdFromUrl(artist.previewUrl));

	let contentScrollPosY = $state(0);
	$inspect(contentScrollPosY);
</script>

<main>
	<div class="grid min-h-svh grid-cols-1 grid-rows-[85svh_1fr]">
		<div class="px-auto relative isolate flex flex-col justify-end bg-blue-500 py-16">
			<img
				src={artist.imageUrl}
				alt="Cover af {artist.name}"
				class="absolute top-0 left-0 h-full w-full object-cover"
			/>
			<Fader direction="right" class="absolute w-96 from-zinc-950" />
			<Fader direction="up" class="absolute h-96 from-zinc-950" />
			<div class="z-10">
				<h1 style:word-spacing="100vw" class="text-7xl font-bold md:text-8xl lg:text-9xl">
					{artist.name}
				</h1>
			</div>
		</div>
		<article class="px-auto space-y-8 bg-zinc-950 py-16">
			<!-- ARTICLE CONTENT -->
			<div class="prose prose-invert max-w-none">
				{@html artist.description}
			</div>

			{#if trackId}
				<SpotifyPreview {trackId} />
			{/if}

			{#if events.length > 0}
				<h1 class="text-2xl font-bold">Kommende events.</h1>
				<EventCaroussel {events} />
			{/if}
		</article>
	</div>
</main>
