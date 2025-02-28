<script lang="ts">
	let { children } = $props();
</script>

<div class="flex justify-end border-t border-t-zinc-800 bg-zinc-900 p-4">
	{@render children()}
</div>

