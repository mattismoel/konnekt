<script lang="ts">
	import type { Genre } from '$lib/features/artist/genre';

	type Props = {
		genres: Genre[];
	};

	let { genres }: Props = $props();
</script>
