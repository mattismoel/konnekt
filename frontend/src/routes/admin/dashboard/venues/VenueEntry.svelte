<script lang="ts">
	import { cn } from '$lib/clsx';
	import Selector from '$lib/components/ui/Selector.svelte';
	import { COUNTRIES_MAP } from '$lib/location';
	import type { Venue, venueForm } from '$lib/venue';
	import type { z } from 'zod';

	import CheckIcon from '~icons/mdi/check';
	import TrashIcon from '~icons/mdi/trash';
	import XIcon from '~icons/mdi/close';
	import Button from '$lib/components/ui/Button.svelte';

	type Props = {
		initialValue?: Venue;
		onEdit: (form: z.infer<typeof venueForm>) => void;
		onDelete: () => void;
	};

	let { initialValue, onEdit, onDelete }: Props = $props();

	let form = $state<z.infer<typeof venueForm>>({
		...(initialValue || {
			name: '',
			city: '',
			countryCode: ''
		})
	});

	let isEdited = $derived(
		form.name !== initialValue?.name ||
			form.city !== initialValue?.city ||
			form.countryCode !== initialValue?.countryCode
	);

	const resetForm = () => {
		form = { ...(initialValue || { name: '', city: '', countryCode: '' }) };
	};
</script>

<svelte:window
	onkeydown={(e) => {
		if (e.key === 'Escape' && isEdited) resetForm();
	}}
/>

<li
	class={cn(
		'group flex items-center gap-8 rounded-md border border-transparent p-2 hover:border-zinc-700 hover:bg-zinc-800',
		{
			'border-green-800 bg-green-950 hover:border-green-800 hover:bg-green-950': isEdited
		}
	)}
>
	<input
		bind:value={form.name}
		class={cn('w-full rounded-sm border-transparent bg-transparent font-semibold', {
			italic: form.name !== initialValue?.name
		})}
	/>
	<input
		bind:value={form.city}
		class={cn('text-text/50 w-full rounded-sm border-transparent bg-transparent', {
			italic: form.city !== initialValue?.city
		})}
	/>
	<Selector
		class={cn('text-text/50 border-transparent bg-transparent group-hover:border-zinc-700', {
			italic: form.countryCode !== initialValue?.countryCode
		})}
		entries={Array.from(COUNTRIES_MAP).map(([key, val]) => ({ name: val, value: key }))}
		bind:value={form.countryCode}
	/>
	<div class="text-text/50 flex w-full justify-end gap-8">
		{#if !isEdited}
			<Button variant="dangerous" onclick={onDelete}>
				<TrashIcon />
			</Button>
		{:else}
			<div class="flex gap-4">
				<button onclick={resetForm} class="hover:text-red-400">
					<XIcon />
				</button>
				<button onclick={() => onEdit(form)} class="hover:text-green-400"><CheckIcon /></button>
			</div>
		{/if}
	</div>
</li>
