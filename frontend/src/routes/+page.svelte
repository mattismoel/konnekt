<script lang="ts">
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import LogoScroller from '$lib/components/ui/LogoScroller.svelte';
	import EventDetails from './events/[id]/EventDetails.svelte';

	import AlbaniLogo from '$lib/assets/AlbaniLogo.svg';
	import MomentumLogo from '$lib/assets/MomentumLogo.svg';
	import OdenseKommuneLogo from '$lib/assets/OdenseKommuneLogo.svg';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<div>
	{#if events.length > 0}
		<EventDetails event={events[0]} />
	{/if}
	<section class="px-auto py-16">
		<h1 class="mb-8 text-2xl font-bold">Kommende events</h1>
		<Caroussel>
			{#each events as event (event.id)}
				<EventCard {event} />
			{/each}
		</Caroussel>
	</section>
	<section class="px-auto space-y-8 bg-zinc-900 py-16">
		<section>
			<h1 class="mb-8 text-2xl font-bold">Vores mission</h1>
			<p class="text-text/75">
				Lorem ipsum dolor sit amet, consectetur adipiscing elit. Etiam tincidunt ultrices magna
				vitae tristique. Morbi ultricies ante molestie ipsum tempus, eget commodo mauris interdum.
				Fusce fringilla nisi vel nisl gravida finibus. Integer fermentum rutrum cursus. Aliquam
				commodo, odio imperdiet mollis tempus, ex velit congue nulla, sit amet lobortis nunc eros id
				ligula. Quisque vel turpis sit amet nisi dictum dignissim. Nunc vestibulum molestie massa a
				vehicula. Praesent auctor tellus quis volutpat semper.
			</p>
		</section>
		<section>
			<h1 class="mb-12 text-2xl font-bold">Med støtte fra</h1>
			<LogoScroller
				srcs={new Map<string, string>([
					['Momentum', MomentumLogo],
					['Albani', AlbaniLogo],
					['Odense Kommune', OdenseKommuneLogo]
				])}
			></LogoScroller>
		</section>
	</section>
</div>
