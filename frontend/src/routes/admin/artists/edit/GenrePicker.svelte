<script lang="ts">
	import type { Genre } from '$lib/artist';

	type Props = {
		genres: Genre[];
	};

	let { genres }: Props = $props();
</script>
