<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Alignment = 'left' | 'center' | 'right';

	type Props = HTMLAttributes<HTMLTableHeaderCellElement> & {
		alignment?: Alignment;
	};

	let { alignment = 'left', children }: Props = $props();
</script>

<th
	class={cn('px-4 py-4 font-medium', {
		'text-left': alignment === 'left',
		'text-center': alignment === 'center',
		'text-right': alignment === 'right'
	})}
>
	{@render children?.()}
</th>
