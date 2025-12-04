<script lang="ts">
	import { page } from "$app/state";

	type Props = {
		entries: { href: string; title: string }[];
		onToggleMenu: () => void;
	};

	const SCROLL_BUFFER = 64;

	let { entries, onToggleMenu }: Props = $props();

	let scrollY = $state(0);
	const scrolled = $derived(scrollY > SCROLL_BUFFER);
</script>

<svelte:window bind:scrollY />

{#snippet entry(href: string, title: string)}
	{@const isCurrent = page.url.pathname === href}
	<li>
		<a
			{href}
			{title}
			class={[
				"transition-colors",
				"before:invisible before:block before:h-0 before:overflow-hidden before:font-semibold before:content-[attr(title)]",
				"hover:text-text-light",
				isCurrent && "font-semibold text-text-light"
			]}
		>
			{title}
		</a>
	</li>
{/snippet}

<nav
	class={["fixed z-40 flex h-nav w-full transition-[inset]", scrolled ? "inset-y-2" : "inset-0"]}
>
	<div
		class={[
			"w-full min-w-fit transition-[max-width,margin]",
			scrolled ? "mx-responsive" : "m-[0_auto] max-w-full"
		]}
	>
		<div
			class={[
				"inset-0 h-nav w-full bg-linear-to-b outline transition-[backdrop-filter,--tw-gradient-from,--tw-gradient-to,padding,border-radius,inset,outline-color] duration-200",
				scrolled
					? "rounded-4xl from-zinc-950/75 to-zinc-950/75 outline-zinc-800 backdrop-blur-2xl"
					: "rounded-none from-black/80 outline-transparent backdrop-blur-none"
			]}
		>
			<div
				class={[
					"inset-0 left-0 flex h-nav items-center justify-between gap-32 transition-[padding] duration-200",
					scrolled ? "px-8 md:px-12" : "px-8"
				]}
			>
				<div class="flex items-center gap-6">
					<a href="/" class="text-xl font-black text-text-light">KONNEKT</a>
					<button type="button" onclick={onToggleMenu} class="sm:hidden">MENU</button>
				</div>

				<ul class="hidden gap-6 sm:flex">
					{#each entries as { href, title }}
						{@render entry(href, title)}
					{/each}
				</ul>
			</div>
		</div>
	</div>
</nav>
