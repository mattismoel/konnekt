<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		selected: boolean;
		name: string;
		onSelect: () => void;
		onDeselect: () => void;
	};

	let { name, selected, onSelect, onDeselect }: Props = $props();
</script>

<button
	type="button"
	onclick={!selected ? onSelect : onDeselect}
	class={cn(
		'text-text flex w-full items-center gap-4 rounded-sm border border-transparent bg-zinc-950 p-2 hover:border-zinc-800',
		{
			'border-zinc-800 bg-zinc-900': selected
		}
	)}
>
	{@render check(selected)}
	{name}
</button>

{#snippet check(selected: boolean)}
	<div class="h-5 w-5 rounded-full border border-zinc-700 bg-zinc-800 p-1">
		<div
			class={cn('h-full w-full rounded-full bg-zinc-700', {
				'bg-blue-500': selected
			})}
		></div>
	</div>
{/snippet}
