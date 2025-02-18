<script>
	let { children } = $props();
	import '../app.css';
</script>

{@render navbar()}
{@render children()}
{#snippet navbar()}
	<nav class="h-nav flex items-center justify-between border-b border-b-neutral-950 px-8">
		<a href="/" class="font-black">KONNEKT&reg;</a>
		<ul class="*:hover:text-text text-text/85 flex gap-6">
			<li><a href="/events">Events</a></li>
			<li><a href="/artists">Kunstnere</a></li>
			<li><a href="/sponsorer">Sponsorer</a></li>
			<li><a href="/om-os">Om os</a></li>
		</ul>
	</nav>
{/snippet}

