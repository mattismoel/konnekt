<script lang="ts">
	import { socialUrlToIcon } from '$lib/features/artist/social';

	import Button from '$lib/components/ui/Button.svelte';
	import TrashIcon from '~icons/mdi/trash';

	type Props = {
		url: string;
		onDelete: () => void;
	};

	let { url = $bindable(''), onDelete }: Props = $props();

	const Icon = socialUrlToIcon(url);
</script>

<div class="flex w-full items-center gap-4">
	<div class="flex w-full items-center gap-4">
		<Icon class="text-lg" />
		<input
			bind:value={url}
			type="text"
			class="w-full rounded-sm border-transparent bg-zinc-950 hover:border-zinc-800"
		/>
	</div>
	<Button type="button" onclick={onDelete} variant="dangerous">
		<TrashIcon />
	</Button>
</div>
