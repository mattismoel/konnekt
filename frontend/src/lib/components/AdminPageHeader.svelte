<script lang="ts">
	import type { WithChildren, WithClass } from "$lib/class";

	type Props = WithChildren &
		WithClass & {
			title: string;
			description: string;
		};

	let { title, description, children, ...rest }: Props = $props();
</script>

<header class={["pb-8", rest.class]}>
	<div class="mb-8">
		<h1 class="font-heading mb-4 text-4xl font-bold text-text-light">{title}</h1>
		<p class="leading-relaxed">{description}</p>
	</div>

	<div class="flex">
		{@render children?.()}
	</div>
</header>
