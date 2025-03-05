<script lang="ts">
	import { artistFormSchema, type Artist, type ArtistForm } from '$lib/artist';
	import Input from '$lib/components/ui/Input.svelte';
	import Pill from '$lib/components/ui/Pill.svelte';
	import PlusIcon from '~icons/mdi/add';
	import type { Genre } from '$lib/genre';
	import GenreSelectorModal from '$lib/components/ui/GenreSelectorModal.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import SocialEntry from './SocialEntry.svelte';
	import type { ZodError } from 'zod';
	import FieldError from '$lib/components/ui/FieldError.svelte';
	import FilePicker from '$lib/components/ui/FilePicker.svelte';
	import ImageSelectorModal from '$lib/components/ui/ImageSelectorModal.svelte';
	import ImagePreview from '$lib/components/ui/ImagePreview.svelte';

	type Props = {
		artist: Artist | null;
		genres: Genre[];
		onSubmit: (form: ArtistForm) => void;
	};

	let { artist, genres, onSubmit }: Props = $props();

	let form: ArtistForm = $state({
		name: artist?.name || '',
		description: artist?.description || '',
		genreIds: artist?.genres.map((genre) => genre.id) || [],
		image: null,
		socials: artist?.socials || []
	});

	let socialUrl = $state('');
	let selectedGenres = $derived(genres.filter((genre) => form.genreIds.includes(genre.id)));

	let showGenreModal = $state(false);
	let formError = $state<ZodError | null>(null);

	let coverImageUrl = $derived(
		form.image ? URL.createObjectURL(form.image) : artist?.imageUrl || ''
	);

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

<form class="w-full max-w-xl space-y-16" onsubmit={submit}>
	<div class="space-y-8">
		<h1 class="mb-8 text-2xl font-bold">Generelt.</h1>
		<ImagePreview src={coverImageUrl} onChange={updateImage} />
		<div class="space-y-8">
			<div class="space-y-1">
				<Input label="Kunstnernavn" bind:value={form.name} />
				<FieldError errors={formError?.flatten().fieldErrors['name']} />
			</div>
			<div class="space-y-1">
				<Input label="Beskrivelse" bind:value={form.description} />
				<FieldError errors={formError?.flatten().fieldErrors['name']} />
			</div>
		</div>
	</div>

	<div>
		<h1 class="mb-4 text-2xl font-bold">Genrer.</h1>
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
					onChange={(newUrl) =>
						(form.socials = form.socials.map((social) => (social === url ? newUrl : social)))}
					onDelete={() => (form.socials = form.socials.filter((social) => social !== url))}
				/>
			{/each}
		</div>
		<FieldError errors={formError?.flatten().fieldErrors['socials']} />
	</div>
	<Button type="submit" expandX>Offentligør</Button>
</form>
