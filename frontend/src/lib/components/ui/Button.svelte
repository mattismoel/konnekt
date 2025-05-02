<script lang="ts">
	import { cn } from '$lib/clsx';
	import type { Snippet } from 'svelte';
	import type { HTMLAnchorAttributes, HTMLButtonAttributes } from 'svelte/elements';

	type Variant = 'outline' | 'primary' | 'secondary' | 'dangerous' | 'ghost';

	type RootProps = {
		variant?: Variant;
		children?: Snippet | undefined;
	};

	type AnchorElement = RootProps &
		Omit<HTMLAnchorAttributes, 'href' | 'type'> & {
			href: HTMLAnchorAttributes['href'];
			type?: never;
			disabled?: HTMLButtonAttributes['disabled'];
		};

	type ButtonElement = RootProps &
		Omit<HTMLButtonAttributes, 'type' | 'href'> & {
			type?: HTMLButtonAttributes['type'];
			href?: never;
			disabled?: HTMLButtonAttributes['disabled'];
		};

	type Props = AnchorElement | ButtonElement;

	let { href, type, disabled = false, variant = 'primary', children, ...rest }: Props = $props();

	const baseClasses =
		'flex h-min w-fit items-center justify-center gap-3 rounded-sm px-3 py-2 font-medium text-zinc-950 transition-colors disabled:opacity-50';

	const variantClasses = new Map<Variant, string>([
		['primary', 'bg-zinc-100 text-zinc-950 hover:bg-zinc-300'],
		[
			'secondary',
			'text-text/85 hover:text-text border border-zinc-700 bg-zinc-800 font-normal hover:border-zinc-600 hover:bg-zinc-700'
		],
		[
			'outline',
			'text-text/85 hover:text-text border border-zinc-100/75 transition-colors hover:border-zinc-100/30 hover:bg-zinc-100/25'
		],
		[
			'dangerous',
			'border border-red-900 bg-red-950 text-red-400 hover:border-red-800 hover:bg-red-900 hover:text-red-300'
		],
		['ghost', 'text-text/75 hover:text-text border border-zinc-900 font-normal hover:bg-zinc-900']
	]);
</script>

<svelte:element
	this={href ? 'a' : 'button'}
	data-button-root
	type={href ? undefined : type}
	href={href && !disabled ? href : undefined}
	disabled={href ? undefined : disabled}
	aria-disabled={href ? disabled : undefined}
	role={href && disabled ? 'link' : undefined}
	tabindex={href && disabled ? -1 : 0}
	class={cn(baseClasses, variantClasses.get(variant), rest.class)}
>
	{@render children?.()}
</svelte:element>
