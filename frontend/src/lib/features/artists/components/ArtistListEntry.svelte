<script lang="ts">
	import type { Artist } from "../artist";

	type Props = {
		artist: Artist;
	};

	let { artist }: Props = $props();
</script>

<li
	class="rounded-2xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800"
>
	<a href="/admin/artists/edit/{artist.id}" class="flex flex-col px-8 py-4">
		<p class="font-medium text-text-light">{artist.name}</p>
		<p class="text-sm">{artist.genres.map((genre) => genre.name).join(", ")}</p>
	</a>
</li>
