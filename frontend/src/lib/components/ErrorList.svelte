<script lang="ts">
	type Props = {
		issues?: string[] | undefined;
	};

	let { issues }: Props = $props();

	$inspect(issues);
</script>

<ul class="flex list-inside list-disc flex-col gap-1 text-sm text-red-500">
	{#each issues as issue}
		<li>{issue}</li>
	{/each}
</ul>
