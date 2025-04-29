<script lang="ts">
	import '../app.css';
	import { page } from '$app/state';
</script>

<main class="bg-background px-auto flex h-svh w-screen items-center justify-center py-16">
	<div class="flex flex-col items-center">
		<h1 class="mb-4 text-center text-7xl font-bold">{page.status}</h1>
		{#if page.status === 401}
			<span>Det ser ud til, at du ikke har tilladelse til at se denne side...</span>
		{:else if page.status === 404}
			<span class="mb-4">Hov! Denne side findes desværre ikke...</span>
			<a
				href="/"
				class="decoration-text/50 hover:decoration-text underline decoration-wavy decoration-2 underline-offset-4 transition-colors duration-300"
				>Lad os følge dig hjem igen...</a
			>
		{:else}
			<span>Øv. Noget gik galt...</span>
		{/if}
	</div>
</main>
