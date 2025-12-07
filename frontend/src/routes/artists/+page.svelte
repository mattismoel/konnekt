<script lang="ts">
	import Button from "$lib/components/ui/Button.svelte";
	import { type Artist } from "$lib/features/artists/artist";
	import { getPreviousArtists, getUpcomingArtists } from "$lib/features/artists/artist.remote";

	const upcomingArtists = $derived(await getUpcomingArtists());
	const previousArtists = $derived(await getPreviousArtists());

	let activeArtist = $derived(upcomingArtists.at(0));

	const handleHover = (id: string) => {
		const artist = upcomingArtists.find((artist) => artist.id === id);
		if (!artist) return;

		activeArtist = artist;
	};

	let showPrevious = $derived(upcomingArtists.length === 0);
</script>

{#snippet entry({ id, name, genres, socials }: Artist)}
	<li
		onmouseover={() => handleHover(id)}
		onfocus={() => handleHover(id)}
		class="grid grid-cols-3 items-center rounded-2xl border border-transparent transition-colors hover:border-foreground/10 hover:bg-foreground/10"
	>
		<a href="/artists/{id}" class="py-2 pl-8 font-medium">
			{name}
		</a>

		<p class="hidden @lg:inline-block">{genres.map((genre) => genre.name).join(", ")}</p>

		<ul>
			{#each socials as { url }}
				<li>
					<a href={url}>
						<p>O</p>
					</a>
				</li>
			{/each}
		</ul>
	</li>
{/snippet}

<main class="min-h-svh">
	{#each upcomingArtists as { id, cover, name }}
		<img
			src={cover}
			alt={name}
			class={[
				"absolute -z-10 h-full w-full object-cover brightness-50 transition-[opacity,scale]",
				activeArtist?.id === id ? "scale-100 opacity-100" : "scale-105 opacity-0"
			]}
		/>
	{/each}

	<div class="z-50 mx-responsive py-32">
		<h1 class="font-heading mb-4 text-4xl font-bold">Kommende kunstnere</h1>

		<ul class="@container">
			{#each upcomingArtists as artist}
				{@render entry(artist)}
			{/each}
		</ul>

		{#if previousArtists.length}
			<Button variant="secondary">Vis tidligere</Button>
			{#if showPrevious}
				<h1>Tidligere kunstnere</h1>

				<ul class="@container">
					{#each previousArtists as artist}
						{@render entry(artist)}
					{/each}
				</ul>
			{/if}
		{/if}
	</div>
</main>
