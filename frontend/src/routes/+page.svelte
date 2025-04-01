<script lang="ts">
	import LandingImage from '$lib/assets/landing.jpg';
	import Caroussel from '$lib/components/Caroussel.svelte';
	import EventCard from '$lib/components/EventCard.svelte';
	import LogoScroller from '$lib/components/ui/LogoScroller.svelte';
	import RightArrowIcon from '~icons/mdi/arrow-right';

	import OdenseKommuneLogo from '$lib/assets/logos/odense-kommune-logo.svg';
	import UngOdenseLogo from '$lib/assets/logos/ungodense-logo.svg';
	import PostenLogo from '$lib/assets/logos/posten-logo.svg';
	import KulturMaskinenLogo from '$lib/assets/logos/kulturmaskinen-logo.svg';
	import SpillestedetOdenseLogo from '$lib/assets/logos/spillestedet-odense-logo.svg';

	import Button from '$lib/components/ui/Button.svelte';
	import Fader from '$lib/components/ui/Fader.svelte';
	import GlowCursor from '$lib/components/GlowCursor.svelte';
	import ArtistDisplay from '$lib/components/ui/ArtistDisplay.svelte';

	let { data } = $props();
	let { events } = $derived(data);
</script>

<div>
	<section class="px-auto -z-50 flex h-svh flex-col justify-center gap-16">
		<div
			class="pointer-events-none absolute top-0 left-0 isolate -z-10 h-full w-full overflow-hidden"
		>
			<GlowCursor />
			<img
				src={LandingImage}
				alt=""
				class="pointer-events-none fixed top-0 left-0 z-0 h-full w-full object-cover brightness-50"
			/>
		</div>
		<section class="flex max-w-lg flex-col gap-8">
			<h2 class="font-heading text-5xl">For et stærkere <b>fynsk musisk vækstlag</b></h2>
			<p class="text-text/75">
				En forening med formål, at støtte det lokale fynske musiske vækstlag og give aspirerende
				musikere et springbræt til den danske musikscene.
			</p>
			<div class="flex w-full flex-col-reverse gap-4 sm:flex-row">
				<form action="/om-os">
					<Button type="submit" variant="secondary" class="h-16 w-full sm:w-fit">Læs mere</Button>
				</form>
				<form action="/events">
					<Button type="submit" class="group h-16 w-full items-center gap-2 sm:w-fit">
						Se events
						<RightArrowIcon class="transition-transform group-hover:translate-x-2" />
					</Button>
				</form>
			</div>
		</section>
		<Fader direction="up" class="absolute h-64 from-black/75" />
	</section>
	<section class="px-auto space-y-16 bg-zinc-950 py-16">
		<!-- ABOUT -->
		<section class="space-y-8 py-16">
			<!-- MISSION STATEMENT -->
			<section>
				<h1 class="font-heading mb-8 text-2xl font-bold">Vores mission</h1>
				<p class="text-text/75">
					Foreningen Konnekt har som formål at støtte unge musikere og skabe en platform, hvor de
					kan vise deres talent frem og få vigtig erfaring med liveoptrædener. Projektet skal gøre
					det nemmere for spirende talenter at finde deres plads i musikmiljøet og opbygge et
					publikum.
					<br />
					<br />
					Samtidig ønsker vi at give publikum – især unge – mulighed for at opdage nye kunstnere i genrer,
					de allerede har stiftet bekendtskab med før. Derudover vil vi styrke musikmiljøet i Odense
					ved at skabe et fællesskab mellem nye og mere erfarne upcoming kunstnere, som kan dele erfaringer,
					inspirere hinanden og måske endda finde samarbejdspartnere. På den måde skaber Konnekt ikke
					kun koncertoplevelser, men også en grobund for kreativ udvikling og vækst i den lokale kultur.
				</p>
			</section>

			<!-- SPONSORS -->
			<section class="z-0 flex w-full flex-col gap-8">
				<span class="font-bold">Med støtte fra</span>
				<div class="relative isolate w-full">
					<Fader direction="right" class="absolute z-50 w-32 from-zinc-950" />
					<Fader direction="left" class="absolute z-50 w-32 from-zinc-950" />
					<LogoScroller
						class="h-10 w-full"
						srcs={new Map<string, string>([
							['Spillestedet Odense', SpillestedetOdenseLogo],
							['UngOdense', UngOdenseLogo],
							['Posten', PostenLogo],
							['Kulturmaskinen', KulturMaskinenLogo],
							['Odense Kommune', OdenseKommuneLogo]
						])}
					/>
				</div>
			</section>
		</section>

		<!-- EVENTS -->
		<section>
			{#if events.length > 0}
				<h1 class="font-heading mb-8 text-2xl font-bold">Kommende events</h1>
				<Caroussel>
					{#each events as event (event.id)}
						<EventCard {event} />
					{/each}
				</Caroussel>
			{/if}
		</section>
	</section>
</div>
