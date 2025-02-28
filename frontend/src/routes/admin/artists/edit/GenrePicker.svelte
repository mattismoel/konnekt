<script lang="ts">
	import type { Genre } from '$lib/genre';

	type Props = {
		genres: Genre[];
	};

	let { genres }: Props = $props();
</script>
