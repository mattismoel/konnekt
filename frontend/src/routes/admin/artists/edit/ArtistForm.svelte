<script lang="ts">
	import type { Artist, ArtistForm } from '$lib/artist';
	import Input from '$lib/components/ui/Input.svelte';
	import Pill from '$lib/components/ui/Pill.svelte';
	import PlusIcon from '~icons/mdi/add';
	import type { Genre } from '$lib/genre';
	import GenreSelectorModal from '$lib/components/ui/GenreSelectorModal.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import SocialEntry from './SocialEntry.svelte';

	type Props = {
		artist: Artist | null;
		genres: Genre[];
	};

	let { artist = $bindable(), genres }: Props = $props();

	let form: ArtistForm = $state({
		name: artist?.name || '',
		description: artist?.description || '',
		genreIds: artist?.genres.map((genre) => genre.id) || [],
		imageUrl: artist?.imageUrl || '',
		socials: artist?.socials || []
	});

	let socialUrl = $state('');
	let selectedGenres = $derived(genres.filter((genre) => form.genreIds.includes(genre.id)));
	let showGenreModal = $state(false);

	const addSocial = () => {
		// Return if already exists.
		if (form.socials.some((social) => social === socialUrl)) return;

		form.socials = [...form.socials, socialUrl];
	};
</script>

<form class="space-y-16">
	<div>
		<h1 class="mb-8 text-2xl font-bold">Generelt.</h1>
		<div class="space-y-8">
			<Input label="Kunstnernavn" bind:value={form.name} />
			<Input label="Beskrivelse" bind:value={form.description} />
		</div>
	</div>

	<div>
		<h1 class="mb-4 text-2xl font-bold">Genrer.</h1>
		<div class="flex gap-2">
			<button
				onclick={() => (showGenreModal = true)}
				class="flex items-center gap-2 rounded-full border border-zinc-900 px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
			>
				<PlusIcon />Tilføj
			</button>
			{#each selectedGenres as genre (genre.id)}
				<Pill>{genre.name}</Pill>
			{/each}
		</div>
	</div>
	<GenreSelectorModal
		{genres}
		show={showGenreModal}
		onClose={() => (showGenreModal = false)}
		onChange={(selected) => (form.genreIds = selected.map((genre) => genre.id))}
	/>

	<div>
		<h1 class="mb-8 text-2xl font-bold">Sociale medier.</h1>
		<div class="mb-4 flex gap-2">
			<Input type="text" label="URL" bind:value={socialUrl} />
			<Button type="button" onclick={addSocial}><PlusIcon />Tilføj</Button>
		</div>
		<div class="space-y-2">
			{#each form.socials as url (url)}
				<SocialEntry
					{url}
					onDelete={() => (form.socials = form.socials.filter((social) => social !== url))}
				/>
			{/each}
		</div>
	</div>
	<Button expandX>Offentligør</Button>
</form>
