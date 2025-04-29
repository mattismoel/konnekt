<script lang="ts">
	let { children } = $props();
</script>

<div
	class="min-h-16 rounded-b-md border-r border-b border-l border-zinc-800 bg-zinc-900 bg-zinc-950"
>
	{@render children()}
</div>
