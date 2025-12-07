<script lang="ts">
	import type { Venue } from "../venue";

	type Props = {
		venue: Venue;
	};

	let { venue }: Props = $props();
</script>

<li
	class="flex rounded-2xl border border-zinc-800 bg-zinc-900 transition-colors hover:border-zinc-700 hover:bg-zinc-800"
>
	<a href="/admin/venues/edit/{venue.id}" class="w-full px-8 py-4">
		<p>{venue.name}</p>
		<p>{venue.city}, {venue.countryCode}</p>
	</a>
</li>
