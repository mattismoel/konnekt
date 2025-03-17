<script lang="ts">
	import { fade } from 'svelte/transition';

	import { cn } from '$lib/clsx';

	import type { Component } from 'svelte';
	import type { Severity, Toast } from '$lib/toaster.svelte';

	import ErrorLogo from '~icons/mdi/error-outline';
	import WarningLogo from '~icons/mdi/warning-outline';
	import InfoLogo from '~icons/mdi/information-outline';
	import CloseIcon from '~icons/mdi/close';

	type Props = {
		toast: Toast;
		onDelete: () => void;
	};

	let { toast, onDelete }: Props = $props();

	let iconMap = new Map<Severity, Component>([
		['warning', WarningLogo],
		['info', InfoLogo],
		['error', ErrorLogo]
	]);

	const Icon = $derived(iconMap.get(toast.severity));
</script>

<div
	transition:fade
	class={cn('relative flex min-w-sm flex-col rounded-sm border p-2', {
		'border-blue-900 bg-blue-950': toast.severity === 'info',
		'border-red-900 bg-red-950': toast.severity === 'error',
		'border-yellow-900 bg-yellow-950': toast.severity === 'warning'
	})}
>
	<div
		class="flex items-center gap-2"
		class:text-blue-200={toast.severity === 'info'}
		class:text-yellow-200={toast.severity === 'warning'}
		class:text-red-200={toast.severity === 'error'}
	>
		<Icon />
		<span class:font-medium={toast.message && toast.message !== ''}>
			{toast.title}
		</span>
		<button
			type="button"
			onclick={onDelete}
			class="absolute top-2 right-2 rounded-full p-1"
			class:hover:bg-red-900={true}><CloseIcon /></button
		>
	</div>
	<span
		class="whitespace-pre-wrap"
		class:text-blue-300={toast.severity === 'info'}
		class:text-yellow-300={toast.severity === 'warning'}
		class:text-red-300={toast.severity === 'error'}
	>
		{toast.message}
	</span>
</div>
