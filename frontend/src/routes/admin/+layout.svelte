<script lang="ts">
	import Toaster from "$lib/components/toaster/Toaster.svelte";
	import { setToastContext, ToasterState } from "../../lib/components/toaster/toaster.svelte.ts";

	let { children } = $props();

	setToastContext(new ToasterState());
</script>

<div>
	{@render children()}

	<Toaster />
</div>
