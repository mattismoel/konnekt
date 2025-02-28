<script lang="ts">
	import CloseIcon from '~icons/mdi/close';

	type Props = {
		label: string;
		onClose: () => void;
	};

	let { label, onClose }: Props = $props();
</script>

<div class="text-text relative flex items-center justify-center bg-zinc-950 py-2">
	<span class="text-sm font-medium">{label}</span>
	<button type="button" onclick={onClose} class="absolute right-4"><CloseIcon /></button>
</div>
