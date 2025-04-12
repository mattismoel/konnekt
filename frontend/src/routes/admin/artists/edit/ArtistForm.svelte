<script lang="ts">
	import type { ZodError } from 'zod';

	import { artistFormSchema, type Artist, type ArtistForm } from '$lib/artist';
	import { trackIdFromUrl } from '$lib/spotify';

	import type { Genre } from '$lib/genre';

	import Input from '$lib/components/ui/Input.svelte';
	import FieldError from '$lib/components/ui/FieldError.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import ImagePreview from '$lib/components/ImagePreview.svelte';
	import Pill from '$lib/components/Pill.svelte';
	import GenreSelectorModal from '$lib/components/GenreSelectorModal.svelte';
	import SpotifyPreview from '$lib/components/SpotifyPreview.svelte';
	import SocialEntry from './SocialEntry.svelte';

	import PlusIcon from '~icons/mdi/add';
	import TipTapEditor from '$lib/components/tiptap/TipTapEditor.svelte';

	type Props = {
		artist: Artist | null;
		genres: Genre[];
		onSubmit: (form: ArtistForm) => void;
	};

	let { artist, genres, onSubmit }: Props = $props();

	let form: ArtistForm = $state({
		name: artist?.name || '',
		description: artist?.description || '',
		previewUrl: artist?.previewUrl || '',
		genreIds: artist?.genres.map((genre) => genre.id) || [],
		image: null,
		socials: artist?.socials || []
	});

	let socialUrl = $state('');
	let selectedGenres = $derived(genres.filter((genre) => form.genreIds.includes(genre.id)));

	let showGenreModal = $state(false);
	let formError = $state<ZodError | null>(null);

	let imageUrl = $derived(form.image ? URL.createObjectURL(form.image) : artist?.imageUrl || '');

	let trackId = $derived(artist?.previewUrl ? trackIdFromUrl(form.previewUrl) : '');

	const updateImage = (file: File | null) => {
		if (!file) return;

		form.image = file;
	};

	const addSocial = () => {
		// Return if already exists.
		if (form.socials.some((social) => social === socialUrl)) return;

		form.socials = [...form.socials, socialUrl];
	};

	const submit = (e: SubmitEvent & { currentTarget: EventTarget & HTMLFormElement }) => {
		e.preventDefault();

		const { data, success, error } = artistFormSchema.safeParse(form);
		if (!success) {
			formError = error;
			return;
		}

		onSubmit(data);
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
				<Input
					label="Kunstnernavn"
					bind:value={form.name}
					errors={formError?.flatten().fieldErrors['name']}
				/>
			</div>
			<div class="space-y-1">
				<TipTapEditor bind:value={form.description} />
				<FieldError errors={formError?.flatten().fieldErrors['description']} />
			</div>
		</div>
	</div>

	<div>
		<h1 class="font-heading mb-4 mb-8 text-2xl font-bold">Genrer</h1>
		<div class="mb-2 flex flex-wrap gap-2">
			<button
				type="button"
				onclick={() => (showGenreModal = true)}
				class="flex items-center gap-2 rounded-full border border-zinc-900 px-4 py-2 hover:border-zinc-800 hover:bg-zinc-900"
			>
				<PlusIcon />Tilføj
			</button>
			{#each selectedGenres as genre (genre.id)}
				<Pill>{genre.name}</Pill>
			{/each}
		</div>
		<FieldError errors={formError?.flatten().fieldErrors['genreIds']} />
		<GenreSelectorModal
			{genres}
			show={showGenreModal}
			onClose={() => (showGenreModal = false)}
			onChange={(selected) => (form.genreIds = selected.map((genre) => genre.id))}
		/>
	</div>

	<div class="flex flex-col">
		<h1 class="font-heading mb-8 text-2xl font-bold">Spotify Preview</h1>
		<div class="space-y-4">
			<Input label="Preview URL" bind:value={form.previewUrl} />
			{#if trackId}
				<SpotifyPreview {trackId} />
			{/if}
		</div>
	</div>

	<div class="flex flex-col">
		<h1 class="font-heading mb-4 text-2xl font-bold">Sociale medier</h1>
		<div class="mb-4 flex w-full gap-2">
			<Input type="text" label="URL" bind:value={socialUrl} class="flex-1" />
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
		<FieldError errors={formError?.flatten().fieldErrors['socials']} />
	</div>
	<Button type="submit">Offentligør</Button>
</form>
