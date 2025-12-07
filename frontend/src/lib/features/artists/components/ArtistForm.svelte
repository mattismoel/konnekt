<script lang="ts">
	import { goto } from "$app/navigation";
	import type { Artist } from "$lib/features/artists/artist";
	import {
		getGenres,
		type createArtist,
		type editArtist
	} from "$lib/features/artists/artist.remote";
	import { getToastContext } from "$lib/components/toaster/toaster.svelte";
	import FormField from "$lib/components/FormField.svelte";
	import ImagePreview from "$lib/components/ImagePreview.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import { hasPermissions } from "$lib/features/auth/auth.remote";

	const MAX_GENRES = 3;

	type EditProps = {
		variant: "edit";
		artist: Artist;
		form: typeof editArtist;
	};

	type CreateProps = {
		variant: "create";
		form: typeof createArtist;
		artist?: never;
	};

	type Props = CreateProps | EditProps;

	let { ...rest }: Props = $props();

	let { cover, name, description, genreIds } = rest.form.fields;
	let coverInput = $state<HTMLInputElement>();
	let socialInput = $state<HTMLInputElement>();

	const genres = await getGenres(undefined);

	const addSocial = () => {
		if (!socialInput) return;

		const socialsCount = rest.form.fields.socials.value()?.length ?? 0;

		rest.form.fields.socials[socialsCount].url.set(socialInput.value);
	};

	const toggleGenre = (id: string) => {
		const prevGenres = genreIds.value() ?? [];

		const isSelected = prevGenres.includes(id);

		if (isSelected) {
			genreIds.set(prevGenres.filter((genreId) => genreId !== id));
			return;
		}

		if (prevGenres.length >= MAX_GENRES) return;

		genreIds.set([...prevGenres, id]);
	};

	$effect(() => {
		if (rest.variant === "create") return;

		rest.form.fields.set({
			...rest.artist,
			artistId: rest.artist.id,
			genreIds: rest.artist.genres.map((genre) => genre.id),
			socials: rest.artist.socials,
			cover: undefined
		});
	});

	const toaster = getToastContext();

	const disabled = $derived(!(await hasPermissions(["artists:edit"])));
</script>

<form
	{...rest.form.enhance(async ({ submit }) => {
		try {
			await submit();
		} catch (e) {
			toaster.add(
				rest.artist ? "Kunne ikke opdatere kunstner" : "Kunne ikke skabe kunstner",
				"dangerous"
			);
			throw e;
		}

		toaster.add(rest.artist ? "Kunstner opdateret" : "Kunstner skabt");
		await goto("/admin/dashboard");
	})}
	enctype="multipart/form-data"
>
	<section class="mb-16">
		{#if rest.variant === "edit"}
			<input {...rest.form.fields.artistId.as("hidden", rest.artist.id)} />
		{/if}

		<FormField issues={cover.issues()?.map((i) => i.message)} class="mb-8">
			<input
				bind:this={coverInput}
				{...rest.form.fields.cover.as("file")}
				class="hidden"
				{disabled}
			/>
			<ImagePreview src={cover.value() ? URL.createObjectURL(cover.value()) : rest.artist?.cover} />
			{#if !disabled}
				<Button type="button" variant="secondary" onclick={() => coverInput?.click()}
					>Vælg...</Button
				>
			{/if}
		</FormField>

		<fieldset class="flex flex-col gap-2">
			<FormField issues={name.issues()?.map((i) => i.message)}>
				<Input {...name.as("text")} placeholder="Kunstnernavn" {disabled} />
			</FormField>

			<FormField issues={description.issues()?.map((i) => i.message)}>
				<Input {...description.as("text")} placeholder="Beskrivelse" {disabled} />
			</FormField>
		</fieldset>
	</section>

	<section class="mb-16">
		<div class="mb-4 flex justify-between">
			<h1 class="text-2xl font-bold">Genrer</h1>
			{#if !disabled}
				<Button href="/admin/genres">+ Tilføj</Button>
			{/if}
		</div>

		<FormField issues={genreIds.issues()?.map((i) => i.message)}>
			{#each genreIds.value(), i}
				<input {...genreIds[i].as("hidden", genreIds[i].value())} />
			{/each}

			<ul class="flex flex-wrap gap-1">
				{#each genres as genre}
					<li>
						<button
							{disabled}
							type="button"
							onclick={() => toggleGenre(genre.id)}
							class={[
								"rounded-full border px-4 py-2 ",
								genreIds.value()?.includes(genre.id)
									? "border-zinc-500 bg-zinc-600"
									: "border-zinc-800 bg-zinc-900"
							]}
						>
							{genre.name}
						</button>
					</li>
				{/each}
			</ul>
		</FormField>
	</section>

	{#if !disabled || (disabled && rest.form.fields.socials.value()?.length)}
		<section class="mb-16">
			<h1 class="mb-8 text-2xl font-bold">Sociale medier</h1>

			<FormField issues={rest.form.fields.socials.issues()?.map((i) => i.message)}>
				{#if !disabled}
					<div class="mb-8 flex gap-4">
						<Input bind:element={socialInput} placeholder="URL" class="w-full" />
						<Button type="button" onclick={addSocial} class="shrink-0">+ Tilføj</Button>
					</div>
				{/if}

				<ul class="flex flex-col gap-2">
					{#each rest.form.fields.socials.value(), i}
						<li class="flex gap-4">
							{#if rest.variant === "edit" && rest.form.fields.socials[i].id.value()}
								<input
									{...rest.form.fields.socials[i].id.as(
										"hidden",
										rest.form.fields.socials[i].id.value()
									)}
								/>
							{/if}
							<Input {...rest.form.fields.socials[i].url.as("url")} class="" />
							<Button variant="dangerous">Slet</Button>
						</li>
					{/each}
				</ul>
			</FormField>
		</section>
	{/if}

	{#if !disabled}
		<div class="flex flex-col gap-2">
			<Button variant="dangerous">Slet</Button>
			<Button class="w-full">Offentligør</Button>
		</div>
	{/if}
</form>
