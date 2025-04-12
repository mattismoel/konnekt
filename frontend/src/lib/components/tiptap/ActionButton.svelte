<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLButtonElement> & {
		active: boolean;
	};

	let { active, children, ...rest }: Props = $props();
</script>

<button
	{...rest}
	type="button"
	class:active
	class={cn('px-3 py-1 text-sm hover:bg-zinc-700', {
		'bg-zinc-500': active
	})}
>
	{@render children?.()}
</button>
