<script lang="ts">
	import type { HTMLButtonAttributes } from 'svelte/elements';
	import MenuIcon from '~icons/mdi/dots-vertical';

	type Props = HTMLButtonAttributes;

	let { ...rest }: Props = $props();
</script>

<button {...rest} type="button" class="rounded-md p-1 hover:bg-zinc-800">
	<MenuIcon />
</button>
