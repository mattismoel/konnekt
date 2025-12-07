<script lang="ts">
	import type { HTMLSelectAttributes } from "svelte/elements";

	type Props = HTMLSelectAttributes;
	let { children, ...rest }: Props = $props();
</script>

<select
	{...rest}
	class={[
		"rounded-full border border-zinc-800 bg-zinc-900 px-8 py-2",
		"disabled:text-text-light/50",
		rest.class
	]}
>
	{@render children?.()}
</select>
