<script lang="ts">
	import AdminPageHeader from "$lib/components/AdminPageHeader.svelte";
	import { getToastContext } from "$lib/components/toaster/toaster.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import {
		createGenre,
		deleteGenre,
		getArtists,
		getGenres,
		getUpcomingArtists
	} from "$lib/features/artists/artist.remote";
	import ArtistListEntry from "$lib/features/artists/components/ArtistListEntry.svelte";
	import { hasPermissions } from "$lib/features/auth/auth.remote";
	import { ClientResponseError } from "pocketbase";

	const upcomingArtists = await getUpcomingArtists();

	const artists = await getArtists(undefined).then(({ items: artists }) =>
		artists.filter((artist) => !upcomingArtists.some((upcoming) => upcoming.id === artist.id))
	);

	$inspect(artists);

	const genres = $derived(await getGenres());

	const toaster = getToastContext();
</script>

<main class="mx-responsive py-32">
	<AdminPageHeader
		title="Kunstnere"
		description="Her har du overblik over alle kunstnere som medvirker og har medvirket til events."
	>
		{#if await hasPermissions(["artists:edit"])}
			<Button href="/admin/artists/create">+ Tilføj</Button>
		{/if}
	</AdminPageHeader>

	<section class="mb-16">
		{#if upcomingArtists.length > 0}
			<div class="mb-8">
				<h2 class="mb-2 font-medium">Kommende kunstnere</h2>
				<ul class="flex flex-col gap-2">
					{#each upcomingArtists as artist}
						<ArtistListEntry {artist} />
					{/each}
				</ul>
			</div>
		{/if}

		{#if artists.length > 0}
			<div class="mb-8">
				<h2 class="mb-2 font-medium">Alle kunstnere</h2>
				<ul class="flex flex-col gap-2">
					{#each artists as artist}
						<ArtistListEntry {artist} />
					{/each}
				</ul>
			</div>
		{/if}
	</section>

	<form
		{...createGenre.enhance(async ({ submit }) => {
			try {
				await submit();
				toaster.add("Genre tilføjet.");
			} catch (e) {
				toaster.add("Kunne ikke tilføje genre", "dangerous");
				throw e;
			}
		})}
		class="mb-8"
	>
		<h1 class="font-heading mb-4 text-2xl font-bold">Genrer</h1>
		{#if await hasPermissions(["artists:edit"])}
			<div class="flex gap-2">
				<Input {...createGenre.fields.name.as("text")} placeholder="Ny genre..." class="w-full" />
				<Button class="shrink-0">+ Tilføj</Button>
			</div>
		{/if}
	</form>

	<ul class="flex gap-2">
		{#each genres as genre}
			{@const deleteGenreForm = deleteGenre.for(genre.id)}

			<li class="rounded-full border border-zinc-800 bg-zinc-900 px-4 py-2">
				<form
					{...deleteGenreForm.enhance(async ({ submit }) => {
						if (!confirm(`Er du sikker på, at du vil slette ${genre.name}?`)) return;
						try {
							await submit();
							toaster.add("Genre slettet");
						} catch (e) {
							toaster.add(
								`Kunne ikke slette genre "${genre.name}".\nEr genren associeret med en kunstner?`,
								"dangerous"
							);
						}
					})}
				>
					<input {...deleteGenreForm.fields.id.as("hidden", genre.id)} />

					<div class="flex gap-4">
						<p>{genre.name}</p>

						{#if await hasPermissions(["artists:edit"])}
							<button>X</button>
						{/if}
					</div>
				</form>
			</li>
		{/each}
	</ul>
</main>
