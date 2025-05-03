<script lang="ts">
	import type { z, ZodError } from 'zod';

	import { createArtistForm, editArtistForm, type Artist } from '$lib/features/artist/artist';
	import { trackIdFromUrl } from '$lib/features/artist/spotify';

	import type { Genre } from '$lib/features/artist/genre';

	import Input from '$lib/components/ui/Input.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import ImagePreview from '$lib/components/ImagePreview.svelte';
	import Pill from '$lib/components/Pill.svelte';
	import GenreSelectorModal from '$lib/components/GenreSelectorModal.svelte';
	import SpotifyPreview from '$lib/components/SpotifyPreview.svelte';
	import SocialEntry from './SocialEntry.svelte';

	import PublishIcon from '~icons/mdi/upload';
	import PlusIcon from '~icons/mdi/add';

	import TipTapEditor from '$lib/components/tiptap/TipTapEditor.svelte';
	import Spinner from '$lib/components/Spinner.svelte';
	import FormField from '$lib/components/ui/FormField.svelte';

	type Props = {
		artist?: Artist;

		genres: Genre[];

		errors:
			| z.typeToFlattenedError<z.infer<typeof createArtistForm> | z.infer<typeof editArtistForm>>
			| undefined;

		loading: boolean;
		onSubmit: (form: z.infer<typeof createArtistForm> | z.infer<typeof editArtistForm>) => void;
	};

	let { artist, genres, loading, onSubmit }: Props = $props();

	let form = $state<z.infer<typeof createArtistForm> | z.infer<typeof editArtistForm>>(
		artist
			? {
					name: artist.name,
					description: artist.description,
					previewUrl: artist.previewUrl || '',
					genreIds: artist.genres.map((genre) => genre.id),
					image: null,
					socials: artist.socials
				}
			: {
					name: '',
					description: '',
					previewUrl: '',
					genreIds: [],
					socials: [],
					image: null
				}
	);

	let socialUrl = $state('');
	let selectedGenres = $derived(genres.filter((genre) => form.genreIds.includes(genre.id)));

	let showGenreModal = $state(false);
	let formError = $state<ZodError | null>(null);

	let imageUrl = $derived(form.image ? URL.createObjectURL(form.image) : artist?.imageUrl || '');

	let trackId = $derived(form.previewUrl ? trackIdFromUrl(form.previewUrl) : undefined);

	const updateImage = (file: File | null) => {
		if (!file) return;

		form.image = file;
	};

	const addSocial = () => {
		// Return if already exists.
		if (form.socials.some((social) => social === socialUrl)) return;

		form.socials = [...form.socials, socialUrl];
		socialUrl = '';
	};

	const submit = (e: SubmitEvent) => {
		e.preventDefault();

		onSubmit(form);
	};
</script>

<form class="w-full max-w-3xl space-y-16" onsubmit={submit}>
	<h1 class="font-heading mb-8 text-4xl font-bold">
		{#if artist}
			Redigér kunstner
		{:else}
			Lav kunstner
		{/if}
	</h1>
	<div class="space-y-8">
		<ImagePreview src={imageUrl} onChange={updateImage} />
		<div class="space-y-8">
			<div class="space-y-1">
				<FormField errors={formError?.flatten().fieldErrors['name']}>
					<Input placeholder="Kunstnernavn" bind:value={form.name} />
				</FormField>
			</div>
			<div class="space-y-1">
				<FormField errors={formError?.flatten().fieldErrors['description']}>
					<TipTapEditor bind:value={form.description} />
				</FormField>
			</div>
		</div>
	</div>

	<div>
		<h1 class="font-heading mb-8 text-2xl font-bold">Genrer</h1>
		<div class="mb-2 flex flex-wrap gap-2">
			<button
				type="button"
				onclick={() => (showGenreModal = true)}
				class="flex items-center gap-2 rounded-full border border-zinc-900 px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
			>
				<PlusIcon />Tilføj
			</button>
			<ul class="flex flex-wrap gap-2">
				{#each selectedGenres as genre (genre.id)}
					<Pill>{genre.name}</Pill>
				{/each}
			</ul>
		</div>
		<FormField errors={formError?.flatten().fieldErrors['genreIds']}>
			<GenreSelectorModal
				{genres}
				bind:show={showGenreModal}
				onChange={(selected) => (form.genreIds = selected.map((genre) => genre.id))}
			/>
		</FormField>
	</div>

	<div class="flex flex-col">
		<h1 class="font-heading mb-8 text-2xl font-bold">Spotify Preview</h1>
		<div class="space-y-4">
			<Input placeholder="Preview-URL" bind:value={form.previewUrl} />
			{#if trackId}
				<SpotifyPreview {trackId} />
			{/if}
		</div>
	</div>

	<div class="flex flex-col">
		<h1 class="font-heading mb-4 text-2xl font-bold">Sociale medier</h1>
		<div class="mb-4 flex w-full gap-2">
			<FormField errors={formError?.flatten().fieldErrors['socials']}>
				<Input type="text" placeholder="URL" bind:value={socialUrl} class="flex-1" />
			</FormField>
			<Button type="button" onclick={addSocial}><PlusIcon />Tilføj</Button>
		</div>
		<div class="space-y-2">
			{#each form.socials as url (url)}
				<SocialEntry
					{url}
					onChange={(newUrl) =>
						(form.socials = form.socials.map((social) => (social === url ? newUrl : social)))}
					onDelete={() => (form.socials = form.socials.filter((social) => social !== url))}
				/>
			{/each}
		</div>
	</div>
	<Button type="submit">
		{#if loading}
			<Spinner />
			Offentligører...
		{:else}
			<PublishIcon />
			Offentligør
		{/if}
	</Button>
</form>
