<script lang="ts">
	import type { Concert } from '$lib/concert';
	import type { Event } from '$lib/event';

	type Props = {
		event: Event | null;
	};

	const { event }: Props = $props();

	const coverImageUrl = $state(event?.coverImageUrl || '');
</script>

<form action="">
	<input type="text" name="title" placeholder="Title" />
	<input type="text" name="description" placeholder="Description" />
	<img src={coverImageUrl} alt="" />

	<button>+ Koncert</button>
	{#each event?.concerts || [] as concert (concert)}
		{@render createConcertCard(concert)}
	{/each}
</form>

{#snippet createConcertCard(concert?: Concert)}
	<div></div>
{/snippet}
