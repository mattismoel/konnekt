<script lang="ts">
	import Input from '$lib/components/ui/Input.svelte';
	import { editMember, type editMemberForm, type Member } from '$lib/features/auth/member';
	import type { Team } from '$lib/features/auth/team';
	import { ZodError, type z } from 'zod';
	import Button from '$lib/components/ui/Button.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';
	import FormField from '$lib/components/ui/FormField.svelte';

	type Props = {
		member: Member;
		currentMember: Member;
		teams: Team[];
	};

	let { member, currentMember }: Props = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editMemberForm>>>();

	let isCurrentMember = $derived(currentMember.id === member.id);

	let form = $state<z.infer<typeof editMemberForm>>({
		email: member.email,
		firstName: member.firstName,
		lastName: member.lastName
	});

	let hasChanged = $derived(
		form.email !== member.email ||
			form.firstName !== member.firstName ||
			form.lastName !== member.lastName
	);

	const handleSubmit = async (e: SubmitEvent) => {
		e.preventDefault();

		try {
			await editMember(fetch, member.id, form);
			toaster.addToast('Medlem opdateret');
			await invalidateAll();
		} catch (e) {
			if (e instanceof ZodError) {
				errors = e.flatten();
			}
			toaster.addToast('Kunne ikke opdatere medlem', 'Noget gik galt', 'error');
		}
	};
</script>

<form onsubmit={handleSubmit} class="flex flex-col gap-8">
	<h1 class="text-2xl font-bold">Generelt.</h1>

	<div class="flex flex-col gap-4">
		<div class="flex gap-4">
			<FormField errors={errors?.fieldErrors.firstName}>
				<Input placeholder="Fornavn" bind:value={form.firstName} disabled={!isCurrentMember} />
			</FormField>
			<FormField errors={errors?.fieldErrors.lastName}>
				<Input placeholder="Efternavn" bind:value={form.lastName} disabled={!isCurrentMember} />
			</FormField>
		</div>

		<FormField errors={errors?.fieldErrors.email}>
			<Input type="email" placeholder="Email" bind:value={form.email} disabled={!isCurrentMember} />
		</FormField>
	</div>

	<Button type="submit" disabled={!hasChanged || !isCurrentMember}>Opdatér</Button>
</form>
