<script lang="ts">
	import { page } from '$app/state';

	import Logo from '$lib/assets/Logo.svelte';

	type Entry = {
		href: string;
		name: string;
	};

	type Props = {
		entries: Entry[];
	};

	const { entries }: Props = $props();

	let scrollY = $state(0);
</script>

<svelte:window onscroll={(e) => (scrollY = e.currentTarget.scrollY)} />

<nav
	class:scrolled={scrollY > 0}
	class="h-nav fixed z-50 flex w-screen items-center justify-between bg-gradient-to-b from-black/80 px-8 outline outline-transparent transition-colors duration-500 [.scrolled]:from-zinc-950 [.scrolled]:to-zinc-950 [.scrolled]:outline-zinc-800"
>
	<a href="/">
		<Logo class="h-4" />
	</a>
	<ul class="hidden items-center gap-6 text-lg text-zinc-50 sm:flex">
		{#each entries as entry}
			{@render navEntry(entry)}
		{/each}
	</ul>
</nav>

{#snippet navEntry({ href, name }: Entry)}
	<li>
		<a
			{href}
			title={name}
			class:is-current={page.url.pathname === href}
			class="hover:text-text text-text/75 [.is-current]:text-text transition-colors before:invisible before:block before:h-0 before:overflow-hidden before:font-medium before:content-[attr(title)] [.is-current]:font-medium"
		>
			{name}
		</a>
	</li>
{/snippet}
