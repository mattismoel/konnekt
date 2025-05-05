<script lang="ts">
	import AvatarImage from '$lib/assets/avatar.png';

	type Props = {
		imageUrl?: string;
		file: File | null;
	};

	let { imageUrl, file = $bindable(null) }: Props = $props();

	let fileInput: HTMLInputElement;

	const onFileChange = (e: Event & { currentTarget: EventTarget & HTMLInputElement }) => {
		if (!e.currentTarget.files) return;

		const newFile = e.currentTarget.files.item(0);

		changeImage(newFile);
	};

	const changeImage = (newFile: File | null) => {
		if (!newFile) {
			file = null;
			imageUrl = '';
			return;
		}

		file = newFile;
		imageUrl = URL.createObjectURL(newFile);
	};
</script>

<div class="relative w-fit">
	<input hidden bind:this={fileInput} type="file" onchange={onFileChange} />
	<img src={imageUrl || AvatarImage} alt="Profile" class="h-28 w-28 rounded-full object-cover" />
	<button
		onclick={() => fileInput.click()}
		type="button"
		class="bg-text absolute right-0 bottom-0 translate-x-1/2 translate-y-1/2 rounded-sm px-2 py-1 text-sm text-zinc-950 shadow-sm"
		>Vælg</button
	>
</div>
