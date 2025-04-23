<script lang="ts">
	import Input from '$lib/components/ui/Input.svelte';
	import { editMember, type editMemberForm, type Member } from '$lib/features/auth/member';
	import type { Team } from '$lib/features/auth/team';
	import { ZodError, type z } from 'zod';
	import Button from '$lib/components/ui/Button.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';

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

<form onsubmit={handleSubmit}>
	<h1 class="mb-8 text-2xl font-bold">Generelt.</h1>
	<div class="flex gap-4">
		<Input
			label="Fornavn"
			bind:value={form.firstName}
			disabled={!isCurrentMember}
			errors={errors?.fieldErrors.firstName}
		/>
		<Input
			label="Efternavn"
			bind:value={form.lastName}
			disabled={!isCurrentMember}
			errors={errors?.fieldErrors.lastName}
		/>
	</div>
	<Input
		type="email"
		label="Email"
		bind:value={form.email}
		disabled={!isCurrentMember}
		errors={errors?.fieldErrors.email}
	/>
	<Button type="submit" disabled={!hasChanged || !isCurrentMember}>Opdatér</Button>
</form>
