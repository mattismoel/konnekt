<script lang="ts">
	let { children } = $props();
</script>

<div class="rounded-full border border-zinc-800 bg-zinc-900 px-3 py-2">
	{@render children()}
</div>
