<script lang="ts">
	import type { Concert } from '$lib/features/concert/concert';
	import { differenceInMinutes, format } from 'date-fns';

	type Props = {
		concert: Concert;
		totalMinutes: number;
		startHour: Date;
	};

	let { concert, totalMinutes, startHour }: Props = $props();

	const concertStartOffset = differenceInMinutes(concert.from, startHour);
	const concertDurationMinutes = differenceInMinutes(concert.to, concert.from);
</script>

<a
	style:top="calc({(concertStartOffset / totalMinutes) * 100}%)"
	style:height="calc({(concertDurationMinutes / totalMinutes) * 100}% - 1px)"
	class="absolute flex w-full justify-between overflow-hidden rounded-sm border border-t border-blue-800 bg-blue-950 p-2 text-sm transition-colors hover:bg-blue-900"
	href="/artists/{concert.artist.id}"
>
	<p class="font-bold text-blue-200">
		{concert.artist.name}
	</p>
	<p class="text-blue-500">
		{format(concert.from, 'HH:mm')} - {format(concert.to, 'HH:mm')}
	</p>
</a>
