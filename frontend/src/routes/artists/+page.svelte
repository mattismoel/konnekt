<script lang="ts">
	import type { Artist } from '$lib/artist.js';
	import { socialUrlToIcon } from '$lib/social.js';

	let { data } = $props();

	let { artists } = $derived(data);

	let selectedArtist = $state<Artist | null>(null);
</script>

<main class="isolate h-svh">
	<div class="px-auto z-50 h-full w-full pt-20">
		<h1 class="mb-8 text-7xl font-bold">Kunstnere.</h1>

		<!-- ARTISTS -->
		<ul class="divide-text/50 max-h-96 divide-y overflow-y-scroll">
			{#each artists as artist (artist.id)}
				{@render entry(artist)}
			{/each}
		</ul>
	</div>
</main>

{#snippet entry(artist: Artist)}
	<li
		class="hover:bg-text group hover:text-zinc-900"
		onmouseenter={(_) => (selectedArtist = artist)}
		onmouseleave={(_) => (selectedArtist = null)}
	>
		<div class="grid w-full grid-cols-3">
			<a href="/artists/{artist.id}" class="col-span-2 grid grid-cols-2 py-3 pl-3">
				<span class="line-clamp-1 font-bold">{artist.name}</span>
				<span class="line-clamp-1">{artist.genres.map((g) => g.name).join(', ')}</span>
			</a>
			<div
				class="text-text/75 flex items-center justify-end gap-2 pr-3 text-lg group-hover:text-zinc-700"
			>
				{#each artist.socials as social}
					{@const Icon = socialUrlToIcon(social)}
					<a href={social} class="hover:text-text group-hover:hover:text-zinc-950"><Icon /></a>
				{/each}
			</div>
		</div>
	</li>
{/snippet}
