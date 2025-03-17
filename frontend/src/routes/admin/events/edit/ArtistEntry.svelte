<script lang="ts">
	import type { Artist } from '$lib/artist';
	import { cn } from '$lib/clsx';

	import OpenIcon from '~icons/mdi/open-in-new';

	type Props = {
		artist: Artist;
		onSelect: () => void;
		onDeselect: () => void;
		selected: boolean;
	};

	let { artist, selected = true, onSelect, onDeselect }: Props = $props();
</script>

<button
	type="button"
	onclick={() => (!selected ? onSelect() : onDeselect())}
	class={cn(
		'flex w-full items-center gap-4 rounded-md border border-transparent p-2 hover:border-zinc-800',
		{
			'border-zinc-800 bg-zinc-900': selected
		}
	)}
>
	<div class="h-5 w-5 rounded-full border border-zinc-700 bg-zinc-800 p-1">
		<div
			class={cn('h-full w-full rounded-full bg-zinc-700', {
				'bg-blue-500': selected
			})}
		></div>
	</div>
	<span class="flex-1 text-left">{artist.name}</span>
	<a class="hover:text-text text-zinc-500" target="_blank" href="/admin/artists/edit?id={artist.id}"
		><OpenIcon /></a
	>
</button>
