<script>
	import '../../app.css';

	import Footer from './Footer.svelte';
	import Navbar from './Navbar.svelte';

	let { children } = $props();
</script>

<Navbar
	entries={[
		{ href: '/events', name: 'Events' },
		{ href: '/artists', name: 'Kunstnere' },
		{ href: '/about', name: 'Om os' }
	]}
/>
{@render children()}
<Footer />
