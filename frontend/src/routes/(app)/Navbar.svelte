<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '$lib/clsx';

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
	class={cn(
		'h-nav border-[] fixed z-50 flex w-screen items-center justify-between bg-gradient-to-b from-black/80 px-8 outline outline-transparent transition-colors duration-500',
		{
			'from-zinc-950 to-zinc-950 outline-zinc-800': scrollY > 0
		}
	)}
>
	<a href="/">
		<Logo class="h-4" />
	</a>
	<ul class="*:hover:text-text hidden items-center gap-6 text-lg text-zinc-50 sm:flex">
		{#each entries as entry}
			{@render navEntry(entry)}
		{/each}
	</ul>
</nav>

{#snippet navEntry({ href, name }: Entry)}
	<li
		class={cn('text-text/75 hover:text-text transition-colors', {
			'text-text font-medium': page.url.pathname === href
		})}
	>
		<a {href}>{name}</a>
	</li>
{/snippet}
