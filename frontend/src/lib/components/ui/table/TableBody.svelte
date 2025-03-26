<script lang="ts">
	let { children } = $props();
</script>

<tbody class="divide-y divide-zinc-900">
	{@render children()}
</tbody>
