<script lang="ts">
	import { socialUrlToIcon } from '$lib/features/artist/social.js';
	import type { Artist } from '$lib/features/artist/artist.js';
	import { pickRandom } from '$lib/array.js';

	let { data } = $props();

	/** @description The rate of which artist auto display changes artist. */
	const AUTO_DISPLAY_RATE = 0.25;

	let { artists } = $derived(data);

	let selectedArtist = $derived<Artist | undefined>(artists.at(0));

	let timeout = $state<NodeJS.Timeout | null>(null);

	const beginAutoDisplay = () => {
		timeout = setInterval(() => {
			const newArtist = pickRandom<Artist>(artists, selectedArtist);

			selectedArtist = newArtist;
		}, 1000 / AUTO_DISPLAY_RATE);
	};

	const endAutoDisplay = () => {
		if (!timeout) return;
		clearInterval(timeout);
	};

	$effect(() => beginAutoDisplay());
</script>

<main class="px-auto h-svh pt-32">
	{#each artists as artist (artist.id)}
		<img
			src={artist.imageUrl}
			alt={artist.name}
			class:opacity-100={selectedArtist?.id === artist.id}
			class:scale-105={selectedArtist?.id === artist.id}
			class="pointer-events-none absolute top-0 left-0 -z-10 h-full w-full object-cover opacity-0 brightness-75 transition-all duration-1000"
		/>
	{/each}
	<div class="space-y-16">
		<section class="flex flex-col">
			<h1 class="font-heading mb-4 text-5xl font-bold md:text-7xl">Kunstnere.</h1>
			<span class="text-text/75">
				Her kan du se alle kunstnere, som medvirker i kommende events.
			</span>
		</section>
		<!-- ARTISTS -->
		{#if artists.length <= 0}
			<span>Der er ingen aktuelle kunstnere i øjeblikket...</span>
		{/if}
		<ul
			class="divide-text/50 max-h-96 divide-y overflow-y-scroll"
			onmouseleave={() => beginAutoDisplay()}
			onmouseenter={() => endAutoDisplay()}
		>
			{#each artists as artist (artist.id)}
				{@render entry(artist)}
			{/each}
		</ul>
	</div>
</main>

{#snippet entry(artist: Artist)}
	<li
		class="group text-text/75 hover:text-text [.selected]:text-text relative flex items-center border-l-transparent transition-colors"
		class:selected={selectedArtist?.id === artist.id}
		onmouseenter={(_) => (selectedArtist = artist)}
	>
		<!-- SELECTED MARKER -->
		<div
			class="group-[.selected]:bg-text h-6 w-1 scale-y-0 rounded-full bg-transparent transition-all group-[.selected]:scale-y-100"
		></div>
		<div class="grid w-full grid-cols-3">
			<a href="/artists/{artist.id}" class="col-span-2 grid grid-cols-2 py-3 pl-3">
				<span class="line-clamp-1 font-bold">{artist.name}</span>
				<span class="line-clamp-1">{artist.genres.map((g) => g.name).join(', ')}</span>
			</a>
			<div
				class="text-text/50 group-[.selected]:text-text/75 group-hover:text-text/75 flex items-center justify-end gap-2 pr-3 text-lg"
			>
				{#each artist.socials as social}
					{@const Icon = socialUrlToIcon(social)}
					<a href={social} class="hover:text-text"><Icon /></a>
				{/each}
			</div>
		</div>
	</li>
{/snippet}
