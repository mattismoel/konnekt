<script lang="ts">
	import type { HTMLInputAttributes } from "svelte/elements";
	import Input from "./ui/Input.svelte";

	type Props = Omit<HTMLInputAttributes, "type"> & {
		search: string;
	};

	let { search = $bindable(""), ...rest }: Props = $props();
</script>

<Input type="search" placeholder="Søg..." bind:value={search} {...rest} />
