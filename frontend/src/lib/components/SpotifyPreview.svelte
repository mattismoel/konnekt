<script lang="ts">
	type Props = {
		trackId: string;
	};

	let { trackId }: Props = $props();
	let src = $derived(
		`https://open.spotify.com/embed/track/${trackId}?utm_source=generator&theme=0`
	);
</script>

<div>
	<iframe
		{src}
		title="Audio preview"
		style:border-radius="12px"
		width="100%"
		height="152"
		frameBorder="0"
		allowfullscreen
		allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
		loading="lazy"
	></iframe>
</div>
