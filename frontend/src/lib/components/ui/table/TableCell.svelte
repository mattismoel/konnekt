<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Alignment = 'left' | 'center' | 'right';

	type Props = HTMLAttributes<HTMLTableHeaderCellElement> & {
		alignment?: Alignment;
	};

	let { alignment = 'left', children, ...rest }: Props = $props();
</script>

<td
	class={cn(
		'px-4 py-2',
		{
			'text-left': alignment === 'left',
			'text-center': alignment === 'center',
			'text-right': alignment === 'right'
		},
		rest.class
	)}
>
	{@render children?.()}
</td>
