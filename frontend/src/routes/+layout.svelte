<script>
	let { children } = $props();
	import '../app.css';
</script>

{@render navbar()}
{@render children()}
{@render footer()}

{#snippet navbar()}
	<nav
		class="h-nav absolute z-50 flex w-screen items-center justify-between bg-gradient-to-b from-black/65 px-8"
	>
		<a href="/" class="font-black">KONNEKT&reg;</a>
		<ul class="*:hover:text-text flex gap-6 text-zinc-50">
			<li><a href="/events">Events</a></li>
			<li><a href="/artists">Kunstnere</a></li>
			<li><a href="/sponsorer">Sponsorer</a></li>
			<li><a href="/om-os">Om os</a></li>
		</ul>
	</nav>
{/snippet}

{#snippet footer()}
	<footer class="px-auto border-t border-t-neutral-900 bg-neutral-950 pt-8 pb-6">
		<div class="mb-8 grid w-full grid-cols-1 sm:grid-cols-2">
			<!-- NAVIGATION -->
			<div class="flex-1">
				<h3 class="mb-2 font-bold">Find rundt.</h3>
				<ul class="text-text/50">
					<li><a href="/">Hjem</a></li>
					<li><a href="/events">Events</a></li>
					<li><a href="/artists">Kunstnere</a></li>
					<li><a href="/om-os">Om os</a></li>
					<li><a href="/sponsorer">Sponsorer</a></li>
				</ul>
			</div>
			<!-- CONTACT INFORMATION -->
			<div class="flex flex-1 flex-col items-end">
				<h3 class="mb-2 font-black">KONNEKT&reg;</h3>
				<ul class="text-text/50 flex flex-col items-end">
					<a href="tel:+4512345678">+45 12 34 56 78</a>
					<a href="mailto:konnekt@mail.dk">konnekt@mail.dk</a>
					<a href="https://google.maps.com/konnekt">Konnektvej 17B, Konnekticut</a>
				</ul>
			</div>
		</div>
		<div class="text-text/50 flex justify-center text-xs">
			<span>&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen&nbsp;KONNEKT&reg;</span>
		</div>
	</footer>
{/snippet}
