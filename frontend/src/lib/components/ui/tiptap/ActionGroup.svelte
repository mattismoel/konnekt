<script lang="ts">
	let { children } = $props();
</script>

<div
	class="flex divide-x divide-zinc-800 overflow-hidden rounded-md border border-zinc-700 bg-zinc-800"
>
	{@render children()}
</div>
