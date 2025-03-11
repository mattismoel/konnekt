<script lang="ts">
	let { data } = $props();

	let { artists } = $derived(data);
</script>

<main class="px-auto py-20">
	<ul>
		{#each artists as artist (artist.id)}
			<li><a href="/artists/{artist.id}">{artist.name}</a></li>
		{/each}
	</ul>
</main>
