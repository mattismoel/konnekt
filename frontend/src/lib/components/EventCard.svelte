<script lang="ts">
	import type { Event } from '$lib/event';
	import { formatDateStr } from '$lib/time';
	//import { formatDateString } from '@/lib/time';
	//import { cn } from '@/lib/utils';

	type Props = {
		event: Event;
		//isLoading: boolean;
	};

	const { event }: Props = $props();
	const earliestConcert = $derived(event.concerts[0]);

	type Pos = {
		x: number;
		y: number;
	};

	let mousePos = $state<Pos>({ x: 0, y: 0 });
	let isFocused = $state(false);

	$inspect(mousePos);

	//const [mousePosX, setMousePosX] = useState(0);
	//const [mousePosY, setMousePosY] = useState(0);
	//if (isLoading) return <Skeleton />
</script>

<a
	class="group"
	href={`/events/${event?.id}`}
	onmouseenter={() => (isFocused = !isFocused)}
	onmouseleave={() => (isFocused = !isFocused)}
>
	<div
		role="none"
		class={'relative h-64 w-full overflow-hidden'}
		onmousemove={(e) => {
			const rect = e.currentTarget.getBoundingClientRect();
			mousePos.x = e.clientX - rect.left;
			mousePos.y = e.clientY - rect.top;
			//setMousePos(e.clientX - rect.left, e.clientY - rect.top);
		}}
	>
		<img
			src={event?.coverImageUrl}
			alt={event?.title}
			class="h-full w-full scale-110 object-cover transition-all duration-200 group-hover:scale-100 group-hover:brightness-100 md:brightness-90"
		/>
		<div
			class="absolute bottom-0 left-0 h-1/2 w-full bg-gradient-to-t from-black/80 opacity-0 transition-opacity duration-300 group-hover:opacity-100"
		></div>
		<div
			class="absolute bottom-0 left-0 h-full w-full border border-white/0 mix-blend-overlay transition-all group-hover:border-white/50"
		></div>
		<div
			class="absolute bottom-0 left-0 flex flex-col px-5 pb-5 text-white transition-all duration-100 md:translate-y-full md:group-hover:translate-y-0"
		>
			<h3 class="mb-2 text-3xl font-bold">{event?.title}</h3>
			<span>{formatDateStr(earliestConcert.from || new Date())}</span>
		</div>
		<div
			class={`pointer-events-none absolute h-72 w-72 -translate-x-1/2 -translate-y-1/2 scale-0 bg-white/50 mix-blend-overlay  blur-3xl transition-transform duration-400 group-hover:scale-100`}
			style:left={`${mousePos.x}px`}
			style:top={`${mousePos.y}px`}
		></div>
	</div>
</a>

{#snippet skeleton()}
	<div
		class="relative h-64 w-full animate-pulse overflow-hidden rounded-md
bg-zinc-900"
	>
		<div class="absolute bottom-0 left-0 flex w-full flex-col px-5 pb-5">
			<div
				class="mb-6 h-8 rounded-md bg-zinc-800"
				style:width="{`calc(100% * ${Math.random()})`}}"
			></div>
			<div class="h-4 w-24 rounded-full bg-zinc-800"></div>
		</div>
	</div>
{/snippet}
