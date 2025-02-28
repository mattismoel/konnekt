<script lang="ts">
	import type { Component } from 'svelte';
	import SpotifyIcon from '~icons/mdi/spotify';
	import InstagramIcon from '~icons/mdi/instagram';
	import ErrorIcon from '~icons/mdi/error';
	import TrashIcon from '~icons/mdi/trash';
	import Button from '$lib/components/ui/Button.svelte';

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

	const SocialIcon = $derived.by(() => {
		const { hostname } = new URL(url);
		const iconUrl = hostname.replace(/^www\./, '');

		return iconMap.get(iconUrl);
	});
</script>

<div class="flex gap-2">
	<div class="relative w-full">
		<input
			type="text"
			value={url}
			onchange={(e) => onChange(e.currentTarget.value)}
			class="w-full rounded-md border-b border-transparent bg-transparent hover:border-zinc-900"
		/>
		<div
			class="absolute top-1/2 right-4 z-50 flex -translate-y-1/2 items-center gap-2 text-zinc-300"
		>
			{#if SocialIcon}
				<SocialIcon />
			{:else}
				<ErrorIcon class="text-red-500" />
			{/if}
		</div>
	</div>
	<Button onclick={() => confirm('Sikker?') && onDelete()} variant="dangerous"><TrashIcon /></Button
	>
</div>
