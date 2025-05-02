<script lang="ts">
	import Logo from '$lib/assets/Logo.svelte';
</script>

<footer class="px-auto border-t border-t-zinc-900 bg-zinc-950 pt-8 pb-6">
	<div class="mb-8 grid w-full grid-cols-1 gap-8 sm:grid-cols-2">
		<!-- NAVIGATION -->
		<div class="flex-1">
			<h3 class="font-heading mb-2 font-bold">Find rundt.</h3>
			<ul class="text-text/50">
				<li><a href="/">Hjem</a></li>
				<li><a href="/events">Events</a></li>
				<li><a href="/artists">Kunstnere</a></li>
				<li><a href="/om-os">Om os</a></li>
			</ul>
		</div>
		<!-- CONTACT INFORMATION -->
		<div class="flex flex-1 flex-col items-start sm:items-end">
			<Logo class="mb-2 hidden h-4 sm:block" />
			<h3 class="font-heading mb-2 font-bold">Kontakt os</h3>
			<address class="not-italic">
				<ul class="text-text/50 flex flex-col items-start sm:items-end">
					<a href="tel:+4512345678">+45 12 34 56 78</a>
					<a href="mailto:konnekt@mail.dk">konnekt@mail.dk</a>
					<a href="https://google.maps.com/konnekt">Konnektvej 17B, Konnekticut</a>
				</ul>
			</address>
		</div>
	</div>
	<div class="text-text/50 flex justify-center text-xs">
		<span>&copy;&nbsp;{new Date().getFullYear()}&nbsp;Foreningen&nbsp;KONNEKT&reg;</span>
	</div>
</footer>
