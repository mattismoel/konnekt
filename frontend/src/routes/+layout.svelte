<script lang="ts">
	import "./layout.css";
	import favicon from "$lib/assets/favicon.svg";
	import Navbar from "$lib/components/Navbar.svelte";

	let { children } = $props();

	let showMenu = $state(false);
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Navbar
	onToggleMenu={() => (showMenu = !showMenu)}
	entries={[
		{ href: "/events", title: "Events" },
		{ href: "/artists", title: "Artists" },
		{ href: "/about", title: "Om os" },
		{ href: "/contact", title: "Kontakt" }
	]}
/>

{@render children()}
