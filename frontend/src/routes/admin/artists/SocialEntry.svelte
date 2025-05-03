<script lang="ts">
	import { socialUrlToIcon } from '$lib/features/artist/social';

	import Button from '$lib/components/ui/Button.svelte';

	import TrashIcon from '~icons/mdi/trash';

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
			class="w-full rounded-sm border-transparent bg-zinc-950 hover:border-zinc-800"
			value={url}
			onchange={(e) => onChange(e.currentTarget.value)}
		/>
	</div>
	<Button onclick={() => confirm('Sikker?') && onDelete()} variant="dangerous">
		<TrashIcon />
	</Button>
</div>
