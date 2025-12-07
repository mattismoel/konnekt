<script lang="ts">
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from "svelte/elements";

	type BaseProps = {
		variant?: "primary" | "secondary" | "dangerous";
	};

	type LinkButtonProps = Omit<HTMLAnchorAttributes, "href" | "type"> & {
		href: string;
	};

	type ButtonProps = HTMLButtonAttributes & {
		href?: never;
	};

	type Props = BaseProps & (LinkButtonProps | ButtonProps);

	let { variant = "primary", children, ...rest }: Props = $props();
</script>

<svelte:element
	this={rest.href ? "a" : "button"}
	{...rest}
	class={[
		"flex items-center justify-center gap-4 rounded-full border px-8 py-2 text-center font-medium transition-colors",
		variant === "primary" && "border-foreground bg-foreground text-text-dark",
		variant === "secondary" &&
			"text-text- border-foreground/10 bg-foreground/10 backdrop-blur-xs hover:border-foreground/30 hover:bg-foreground/20 hover:text-text-dark",
		variant === "dangerous" && "border-red-900 bg-red-950 text-red-400",
		rest.class
	]}
>
	{@render children?.()}
</svelte:element>
