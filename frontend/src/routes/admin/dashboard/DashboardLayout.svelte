<script lang="ts">
	let { children } = $props();
</script>

<main class="space-y-16 px-8 py-16 md:px-16">
	{@render children()}
</main>
