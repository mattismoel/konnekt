<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLLiAttributes } from 'svelte/elements';

	type Size = 'sm' | 'base';

	type Props = HTMLLiAttributes & {
		size?: Size;
	};

	let { size = 'sm', children, ...rest }: Props = $props();
</script>

<li
	{...rest}
	class={cn('h-min w-fit rounded-full border border-zinc-800 bg-zinc-900 px-3 py-2', {
		'text-sm': size === 'sm'
	})}
>
	{@render children?.()}
</li>
