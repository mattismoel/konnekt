<script lang="ts">
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { error } from '@sveltejs/kit';
	import Button from './Button.svelte';
	import FilePicker from './FilePicker.svelte';
	import Modal from './modal/Modal.svelte';
	import ModalContent from './modal/ModalContent.svelte';
	import ModalFooter from './modal/ModalFooter.svelte';
	import ModalHeader from './modal/ModalHeader.svelte';

	type Props = {
		label: string;
		show: boolean;
		url: string;
		onClose: () => void;
		onSelect: (url: string) => void;
	};

	let { label, show, url, onClose, onSelect }: Props = $props();

	let file: File | null = $state(null);

	const uploadImage = async () => {
		if (!file) return;

		const formData = new FormData();

		formData.append('file', file);

		const res = await fetch(`${PUBLIC_BACKEND_URL}/artists/coverImage`, {
			method: 'POST',
			body: formData
		});

		if (!res.ok) {
			return error(500, 'Could not update cover image');
		}
	};
</script>

<Modal {show} class="max-w-xl">
	<ModalHeader {label} {onClose} />
	<ModalContent class="max-h-96 p-0">
		<img src={url} alt="Billede" class="w-full" />
	</ModalContent>
	<ModalFooter class="justify-between">
		<FilePicker onChange={(files) => (file = files?.[0] || null)} />
		<Button onclick={uploadImage}>Vælg</Button>
	</ModalFooter>
</Modal>
