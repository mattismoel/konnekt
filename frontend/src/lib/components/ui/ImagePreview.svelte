<script lang="ts">
	import type { HTMLImgAttributes } from 'svelte/elements';
	import FilePicker from './FilePicker.svelte';

	type Props = HTMLImgAttributes & {
		onChange: (file: File) => void;
	};

	let { src, onChange, ...rest }: Props = $props();

	let url = $state(src);

	const updateImage = (file: File | null) => {
		if (!file) return;

		url = URL.createObjectURL(file);
		onChange(file);
	};
</script>

<div class="relative aspect-video w-full">
	<div class="absolute h-full w-full bg-gradient-to-t from-black/50"></div>
	<img {...rest} src={url} class="h-full w-full rounded-md border border-zinc-900 object-cover" />
	<FilePicker
		onChange={(files) => updateImage(files?.[0] || null)}
		class="absolute bottom-4 left-4"
	/>
</div>
