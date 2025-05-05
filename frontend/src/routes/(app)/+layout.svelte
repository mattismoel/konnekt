<script lang="ts">
	import '../../app.css';

	import Footer from './Footer.svelte';
	import * as Navbar from '$lib/components/navbar/index';
	import Logo from '$lib/assets/Logo.svelte';
	import MenuIcon from '~icons/mdi/menu';

	import * as NavMenu from '$lib/components/nav-menu/index';
	import { beforeNavigate, onNavigate } from '$app/navigation';

	let { children } = $props();

	let navMenuExpanded = $state(false);

	beforeNavigate(() => (navMenuExpanded = false));
</script>

<Navbar.Root>
	<Navbar.Header>
		<button>
			<MenuIcon class="text-xl" onclick={() => (navMenuExpanded = true)} />
		</button>
		<a href="/">
			<Logo class="h-4" />
		</a>
	</Navbar.Header>

	<Navbar.RouteList>
		<Navbar.RouteEntry pathname="/events" name="Events" />
		<Navbar.RouteEntry pathname="/artists" name="Kunstnere" />
		<Navbar.RouteEntry pathname="/om-os" name="Om os" />
	</Navbar.RouteList>
</Navbar.Root>

<NavMenu.Root bind:show={navMenuExpanded}>
	<NavMenu.RouteList>
		<NavMenu.Route href="/">Forside</NavMenu.Route>
		<NavMenu.Route href="/events">Events</NavMenu.Route>
		<NavMenu.Route href="/artists">Kunstnere</NavMenu.Route>
		<NavMenu.Route href="/about">Om os</NavMenu.Route>
	</NavMenu.RouteList>
</NavMenu.Root>

{@render children()}

<Footer />
