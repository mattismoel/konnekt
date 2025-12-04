<script lang="ts">
	import type { WithClass } from "$lib/class";
	import type { Attachment } from "svelte/attachments";

	let mousePos = $state<{ x: number; y: number }>({ x: 0, y: 0 });

	const moveHandler: Attachment<Window> = (window) => {
		let rawMousePos = $state<{ x: number; y: number }>({ x: 0, y: 0 });

		const handleMouseMove = (e: MouseEvent) => {
			rawMousePos = { x: e.clientX, y: e.clientY };
			update();
		};

		const update = () => {
			mousePos = {
				x: rawMousePos.x + window.scrollX,
				y: rawMousePos.y + window.scrollY
			};
		};

		$effect(() => {});

		window.addEventListener("mousemove", handleMouseMove);
		window.addEventListener("scroll", update);

		return () => window.removeEventListener("mousemove", handleMouseMove);
	};

	let { ...rest }: WithClass = $props();
</script>

<svelte:window {@attach moveHandler} />

<div
	style:left="{mousePos.x}px"
	style:top="{mousePos.y}px"
	class={[
		"pointer-events-none absolute hidden h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-white/20 mix-blend-overlay blur-3xl brightness-100 sm:block",
		rest.class
	]}
>
	Hello
</div>
