<script lang="ts">
	import type { HTMLInputAttributes } from "svelte/elements";

	type Props = HTMLInputAttributes & {
		variant?: "lighter" | "darker";
		element?: HTMLInputElement;
	};

	let {
		variant = "lighter",
		element = $bindable(),
		value = $bindable(),
		...rest
	}: Props = $props();
</script>

<input
	{...rest}
	bind:this={element}
	bind:value
	class={[
		"rounded-full border  px-8 py-2",
		"disabled:text-text-light/50",
		variant === "lighter" && "border-zinc-800 bg-zinc-900",
		variant === "darker" && "border-zinc-900 bg-zinc-950",
		rest.class
	]}
/>
