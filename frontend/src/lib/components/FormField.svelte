<script lang="ts">
	import type { WithChildren, WithClass } from "$lib/class";
	import ErrorList from "./ErrorList.svelte";

	type Props = (WithClass & WithChildren) & {
		issues?: string[] | undefined;
	};

	let { issues, children, ...rest }: Props = $props();
</script>

<div class={["flex w-full flex-col gap-2", rest.class]}>
	{@render children?.()}

	{#if issues && issues.length > 0}
		<ErrorList {issues} />
	{/if}
</div>
