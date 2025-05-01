<script lang="ts">
	import { cn } from '$lib/clsx';

	type Props = {
		errors?: string[];
		maxCount?: number;
	};

	let { errors, maxCount = 5 }: Props = $props();
</script>

<ul
	class={cn('hidden min-h-5 list-inside text-sm text-red-500', {
		block: errors && errors.length > 0,
		'list-disc': errors && errors.length > 1
	})}
>
	{#each errors?.slice(0, maxCount) || [] as error}
		<li>{error}</li>
	{/each}
</ul>
