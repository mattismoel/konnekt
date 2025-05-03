<script lang="ts">
	import SocialEntry from './SocialEntry.svelte';

	type Props = {
		socials: string[];
	};

	let { socials = $bindable([]) }: Props = $props();

	const handleDeleteSocial = (url: string) => {
		const social = socials.find((social) => social === url);

		if (!social) return;

		if (!confirm('Er du sikker på, at du vil slette linket?')) return;

		socials = socials.filter((social) => social !== url);
	};
</script>

<div class="flex flex-col gap-2">
	{#each socials as url, i (url)}
		<SocialEntry
			bind:url={() => socials[i], (url) => (socials[i] = url)}
			onDelete={() => handleDeleteSocial(url)}
		/>
	{/each}
</div>
