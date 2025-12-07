<script lang="ts">
	import type { Artist } from "$lib/features/artist";
	import { createEvent, editEvent } from "$lib/features/event.remote";
	import Card from "./Card.svelte";
	import FormField from "./FormField.svelte";
	import Input from "./ui/Input.svelte";
	import Select from "./ui/Select.svelte";

	type Props = {
		idx: number;
		artists: Artist[];
		isEdit: boolean;
	};

	let { artists, idx, isEdit }: Props = $props();

	const form = $derived(isEdit ? createEvent : editEvent);
</script>
