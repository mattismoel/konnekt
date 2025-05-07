<script lang="ts">
	import { socialUrlToIcon } from '$lib/features/artist/social';

	import Button from '$lib/components/ui/Button.svelte';
	import TrashIcon from '~icons/mdi/trash';
	import Input from '$lib/components/ui/Input.svelte';

	type Props = {
		url: string;
		onDelete: () => void;
		disabled?: boolean;
	};

	let { url = $bindable(''), disabled = false, onDelete }: Props = $props();

	const Icon = socialUrlToIcon(url);
</script>

<div class:disabled class="[.disabled]:text-text/50 flex w-full items-center gap-4">
	<div class="flex w-full items-center gap-4">
		<Icon class="text-lg" />
		<Input
			{disabled}
			bind:value={url}
			type="text"
			class="w-full rounded-sm border-transparent bg-zinc-950 hover:border-zinc-800"
		/>
	</div>
	{#if !disabled}
		<Button type="button" onclick={onDelete} variant="dangerous">
			<TrashIcon />
		</Button>
	{/if}
</div>
