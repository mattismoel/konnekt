<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLInputElement> & {
		onChange: (files: FileList | null) => void;
	};

	let { onChange, ...rest }: Props = $props();
</script>

<input
	{...rest}
	type="file"
	class={cn(
		'text-text file:rounded-md file:bg-zinc-100 file:px-4 file:py-2 file:font-medium file:text-zinc-900',
		rest.class
	)}
	onchange={(e) => onChange(e.currentTarget.files)}
/>
