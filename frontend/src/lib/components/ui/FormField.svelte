<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		errors?: string[] | undefined;
		maxErrorCount?: number;
	};

	let { children, maxErrorCount = 5, errors }: Props = $props();
</script>

<div class="flex w-full flex-col gap-2">
	{@render children?.()}

	<ul
		class={cn('hidden min-h-5 list-inside text-sm text-red-500', {
			block: errors && errors.length > 0,
			'list-disc': errors && errors.length > 1
		})}
	>
		{#each errors?.slice(0, maxErrorCount) || [] as error}
			<li>{error}</li>
		{/each}
	</ul>
</div>
