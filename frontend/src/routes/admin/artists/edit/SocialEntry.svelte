<script lang="ts">
	import type { Component } from 'svelte';

	import { socialUrlToIcon } from '$lib/features/artist/social';

	import Button from '$lib/components/ui/Button.svelte';

	import SpotifyIcon from '~icons/mdi/spotify';
	import InstagramIcon from '~icons/mdi/instagram';
	import TrashIcon from '~icons/mdi/trash';

	const iconMap = new Map<string, Component>([
		['spotify.com', SpotifyIcon],
		['instagram.com', InstagramIcon]
	]);

	type Props = {
		url: string;
		onChange: (newValue: string) => void;
		onDelete: () => void;
	};

	let { url, onChange, onDelete }: Props = $props();

	const Icon = socialUrlToIcon(url);
</script>

<div class="flex w-full items-center gap-4">
	<div class="flex w-full items-center gap-4">
		<Icon class="text-lg" />
		<input
			type="text"
			class="flex-1 rounded-sm border-transparent bg-zinc-950 hover:border-zinc-800"
			value={url}
			onchange={(e) => onChange(e.currentTarget.value)}
		/>
	</div>
	<Button onclick={() => confirm('Sikker?') && onDelete()} variant="dangerous">
		<TrashIcon />
	</Button>
</div>
