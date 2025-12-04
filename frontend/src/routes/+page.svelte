<script lang="ts">
	import OdenseKommuneLogo from "$lib/assets/logos/odense-kommune-logo.svg";
	import UngOdenseLogo from "$lib/assets/logos/ungodense-logo.svg";
	import PostenLogo from "$lib/assets/logos/posten-logo.svg";
	import KulturMaskinenLogo from "$lib/assets/logos/kulturmaskinen-logo.svg";
	import SpillestedetOdenseLogo from "$lib/assets/logos/spillestedet-odense-logo.svg";
	import StormsPakhusLogo from "$lib/assets/logos/storms-pakhus-logo.svg";
	import NoctivagaLogo from "$lib/assets/logos/noctivaga-logo.svg";
	import GlowCursor from "$lib/components/GlowCursor.svelte";
	import { Randomiser } from "$lib/random.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import SponsorDisplay from "$lib/components/SponsorDisplay.svelte";
	import EventGrid from "$lib/components/EventGrid.svelte";

	let { data } = $props();

	const randomiser = new Randomiser(data.landingImages);

	$effect(() => {
		const interval = setInterval(() => {
			randomiser.randomise();
		}, 3000);

		return () => clearInterval(interval);
	});
</script>

<svelte:head>
	<title>Konnekt | Forside</title>
	<meta title="description" content="Her kan du finde kommende events, aktuelle kunstnere mm." />
</svelte:head>

<section class="-z-50 mx-responsive flex h-svh flex-col justify-center gap-16">
	<div
		class="pointer-events-none absolute top-0 left-0 isolate -z-10 h-full w-full overflow-hidden"
	>
		<div class="absolute h-full w-full">
			<GlowCursor class="z-10" />
			{#each data.landingImages as { image: src, id }, i}
				<img
					{src}
					alt="Baggrundsbillede {i}"
					class={[
						"absolute z-0 h-full w-full object-cover brightness-70 transition-[opacity,scale] duration-500",
						randomiser.entry?.id === id ? "scale-100 opacity-100" : "scale-102 opacity-0"
					]}
				/>
			{/each}
		</div>
	</div>
	<section class="flex max-w-lg flex-col gap-16 overflow-hidden">
		<div class="flex flex-col">
			<h2
				class="font-heading mb-6 text-4xl leading-[1.2] text-text-light text-shadow-lg/15 sm:text-5xl"
			>
				<b>Fynsk musik</b> med fremtiden for øje
			</h2>
			<p class="text-text/75 leading-relaxed text-shadow-md">
				Et springbræt for unge, aspirerende musiskere, og en indgang ind til den danske musikscene.
			</p>
		</div>

		<div class="z-10 flex w-full flex-col-reverse gap-2 sm:flex-row">
			<Button href="/about" variant="secondary" class="w-full sm:w-fit">Læs mere</Button>
			<Button
				href={data.upcomingEvents.length > 0 ? "/events" : "/artists"}
				class="group w-full items-center gap-2 sm:w-fit"
			>
				{#if data.upcomingEvents.length > 0}
					Se events
				{:else}
					Se kunstnere
				{/if}
				<!-- <FaArrowRight class="text-sm transition-transform group-hover:translate-x-1" /> -->
			</Button>
		</div>
	</section>
	<!-- <Fader direction="up" class="absolute z-0 h-64 from-black/75" /> -->
</section>

<section class="border-t border-t-zinc-900 bg-zinc-950">
	<section class="mx-responsive flex flex-col gap-32 py-16">
		<section>
			<h1 class="font-heading mb-8 text-2xl font-bold">Vores mission</h1>
			<p class="text-text/75 leading-loose">
				Konnekt er en ungedrevet forening og et koncertinitiativ, der arbejder for at give unge
				musikere mulighed for at komme på scenen og få kontakt til et publikum. Vi arrangerer
				koncerter, hvor nye artister spiller sammen med mere etablerede upcoming-navne, og hvor
				publikum får adgang til musik, de ellers ikke ville have mødt.
				<br />
				<br />
				Men Konnekt handler ikke kun om musik – det handler om fællesskab og deltagelse. Vi skaber rum
				for unge, der vil være med som frivillige, arrangører eller idéfolk, og vi samarbejder med lokale
				aktører for at skabe synlighed, netværk og udvikling i Odenses kulturliv.
				<br />
				<br />
				Hos os er alle med til at forme oplevelsen – både på og bag scenen.
			</p>
		</section>

		<section class="z-0 flex w-full flex-col gap-8">
			<span class="text-text/50 text-center">I samarbejde med</span>
			<div class="relative isolate w-full">
				<SponsorDisplay
					srcs={[
						{ href: "https://postenlive.dk/", src: PostenLogo, title: "Posten" },
						{ href: "https://stormspakhus.dk/", src: StormsPakhusLogo, title: "Storms Pakhus" },
						{ href: "https://ungodense.dk/", src: UngOdenseLogo, title: "UngOdense" },
						{ href: "https://odense.dk", src: OdenseKommuneLogo, title: "Odense Kommune" },
						{
							href: "https://ungodense.dk/index.php?open=1283&menu_id=58",
							src: SpillestedetOdenseLogo,
							title: "Spillestedet"
						},
						{
							href: "https://kulturmaskinen.dk/",
							src: KulturMaskinenLogo,
							title: "Kulturmaskinen"
						},
						{ href: "https://www.noctivaga.dk/", src: NoctivagaLogo, title: "Noctivaga" }
					]}
				/>
			</div>
		</section>
		<section>
			<h1 class="font-heading mb-16 text-center text-4xl font-bold">Mød holdet</h1>
			<!-- <TeamDisplay allTeams={teams} {members} /> -->
		</section>
		{#if data.upcomingEvents.length > 0}
			<section>
				<h1 class="font-heading mb-8 text-4xl font-bold">Ses vi her?</h1>
				<EventGrid events={data.upcomingEvents} />
			</section>
		{/if}
	</section>
</section>
