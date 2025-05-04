<script lang="ts">
	import type { WithChildren } from '$lib/types';

	type Props = WithChildren<{}>;

	let { children }: Props = $props();

	let scrollY = $state(0);
</script>

<svelte:window onscroll={(e) => (scrollY = e.currentTarget.scrollY)} />

<nav
	class:scrolled={scrollY > 0}
	class="h-nav px-auto fixed z-40 flex w-screen items-center justify-between bg-gradient-to-b from-black/80 outline outline-transparent transition-colors duration-500 [.scrolled]:from-zinc-950 [.scrolled]:to-zinc-950 [.scrolled]:outline-zinc-800"
>
	{@render children?.()}
</nav>
