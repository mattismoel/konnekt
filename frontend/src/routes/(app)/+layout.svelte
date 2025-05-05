<script>
	import '../../app.css';

	import Footer from './Footer.svelte';
	import * as Navbar from '$lib/components/navbar/index';
	import Logo from '$lib/assets/Logo.svelte';
	import MenuIcon from '~icons/mdi/menu';

	let { children } = $props();
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

{@render children()}

<Footer />
