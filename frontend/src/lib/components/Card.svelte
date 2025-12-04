<script lang="ts">
	import type { HTMLAttributes } from "svelte/elements";

	type Props = Omit<HTMLAttributes<HTMLDivElement>, "title"> & {
		title: string;
		description?: string;
		variant?: "darker" | "lighter";
	};

	let { title, description, variant = "darker", children, ...rest }: Props = $props();
</script>

<div
	class={[
		"rounded-xl border px-8 py-12",
		variant === "darker" && "border-zinc-900 bg-zinc-950",
		variant === "lighter" && "border-zinc-800 bg-zinc-900",
		,
		rest.class
	]}
>
	<div class="mb-8">
		<h1 class="font-heading mb-4 text-4xl font-black text-text-light">{title}</h1>
		{#if description}
			<p class="leading-relaxed">{description}</p>
		{/if}
	</div>

	{@render children?.()}
</div>
