<script lang="ts">
	import type { z } from 'zod';
	import ArtistForm from './ArtistForm.svelte';
	import type { artistFormSchema } from '$lib/artist';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { error } from '@sveltejs/kit';

	let { data } = $props();

	const submit = async (form: z.infer<typeof artistFormSchema>) => {
		const res = await fetch(`${PUBLIC_BACKEND_URL}/artists`, {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(form)
		});

		if (!res.ok) {
			return error(500, 'Could not create artist');
		}
	};
</script>

<main class="flex min-h-svh justify-center py-20">
	<ArtistForm artist={data.artist} genres={data.genres} onSubmit={submit} />
</main>
