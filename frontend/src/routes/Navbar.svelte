<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '$lib/clsx';

	import Logo from '$lib/assets/Logo.svelte';

	const entries = new Map<string, string>([
		['/events', 'Events'],
		['/artists', 'Kunstnere'],
		['/om-os', 'Om os']
	]);

	let scrollY = $state(0);
</script>

<svelte:window onscroll={(e) => (scrollY = e.currentTarget.scrollY)} />
<nav
	class={cn(
		'h-nav border-[] fixed z-50 flex w-screen items-center justify-between border-b bg-gradient-to-b from-black/80 px-8 transition-colors duration-500',
		{
			'border-solid border-zinc-900 from-zinc-950 to-zinc-950': scrollY > 0
		}
	)}
>
	<a href="/">
		<Logo class="h-4" />
	</a>
	<ul class="*:hover:text-text hidden items-center gap-6 text-lg text-zinc-50 sm:flex">
		{#each [...entries] as [pathname, name]}
			{@render navEntry(pathname, name)}
		{/each}
	</ul>
</nav>

{#snippet navEntry(pathname: string, name: string)}
	<li
		class={cn('text-text/75 hover:text-text transition-colors', {
			'text-text font-medium': page.url.pathname === pathname
		})}
	>
		<a href={pathname}>{name}</a>
	</li>
{/snippet}
