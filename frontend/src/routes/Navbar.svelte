<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '$lib/clsx';

	import Logo from '$lib/assets/Logo.svelte';

	const entries = new Map<string, string>([
		['/events', 'Events'],
		['/artists', 'Kunstnere'],
		['/sponsorer', 'Sponsorer'],
		['/om-os', 'Om os']
	]);
</script>

<nav
	class="h-nav absolute z-50 flex w-screen items-center justify-between bg-gradient-to-b from-black/80 px-8"
>
	<a href="/">
		<Logo class="h-3" />
	</a>
	<ul class="*:hover:text-text hidden items-center gap-6 text-zinc-50 sm:flex">
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
