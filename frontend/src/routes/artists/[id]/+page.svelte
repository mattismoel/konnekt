<script lang="ts">
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
		<div class="px-auto relative isolate flex items-end py-16">
			<img
				src={artist.imageUrl}
				alt="Cover af {artist.name}"
				class="absolute top-0 left-0 h-full w-full object-cover"
			/>
			<Fader direction="right" class="absolute hidden w-96 from-zinc-950 md:block" />
			<Fader direction="up" class="absolute h-[512px] from-zinc-950" />
			<div
				class="z-10 flex w-full flex-col items-start justify-between gap-8 md:flex-row md:items-end"
			>
				<h1 style:word-spacing="100vw" class="text-7xl font-bold md:text-8xl lg:text-9xl">
					{artist.name}
				</h1>
				<div class="text-text/50 flex gap-4 text-3xl">
					{#each artist.socials as social}
						{@const Icon = socialUrlToIcon(social)}
						<a href={social} class="hover:text-text transition-colors">
							<Icon />
						</a>
					{/each}
				</div>
			</div>
		</div>
		<article class="px-auto space-y-8 bg-zinc-950 py-16">
			<!-- ARTICLE CONTENT -->
			<div class="prose prose-lg md:prose-base prose-invert max-w-none">
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
