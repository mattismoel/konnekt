<script lang="ts">
	import type { HTMLAttributes } from 'svelte/elements';

	type Props = HTMLAttributes<HTMLDivElement> & {
		title: string;
		description: string;
	};

	let { title, description, children }: Props = $props();
</script>

<header class="flex flex-col gap-8">
	<div>
		<h1 class="mb-4 text-4xl font-bold">{title}</h1>
		<p class="text-text/50">{description}</p>
	</div>
	{@render children?.()}
</header>
