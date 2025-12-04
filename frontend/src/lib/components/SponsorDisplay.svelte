<script lang="ts">
	type Props = {
		srcs: { src: string; href: string; title: string }[];
	};

	let { srcs }: Props = $props();
</script>

<div
	class="flex flex-wrap justify-center gap-12 bg-radial from-zinc-900 to-background to-75% p-8 sm:p-16"
>
	{#each srcs as { title, src, href }}
		<a aria-label={title} {href} class="group flex w-20 items-center justify-center sm:w-24">
			<img
				{src}
				alt={title}
				loading="lazy"
				class="aspect-square transition-[filter,scale] group-hover:brightness-100 md:brightness-60"
			/>
		</a>
	{/each}
</div>
