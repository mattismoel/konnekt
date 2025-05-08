<script lang="ts">
	import Input from '$lib/components/ui/Input.svelte';
	import { editMember, type editMemberForm, type Member } from '$lib/features/auth/member';
	import type { Team } from '$lib/features/auth/team';
	import { ZodError, type z } from 'zod';
	import Button from '$lib/components/ui/Button.svelte';
	import { toaster } from '$lib/toaster.svelte';
	import { invalidateAll } from '$app/navigation';
	import FormField from '$lib/components/ui/FormField.svelte';
	import { authStore } from '$lib/auth.svelte';
	import ProfilePictureSelector from '$lib/components/ProfilePictureSelector.svelte';
	import MemberStatusIndicator from '$lib/components/MemberStatusIndicator.svelte';

	type Props = {
		member: Member;
		teams: Team[];
	};

	let { member, teams }: Props = $props();

	let errors = $state<z.typeToFlattenedError<z.infer<typeof editMemberForm>>>();

	let isCurrentMember = $derived(authStore.member?.id === member.id);

	let form = $state<z.infer<typeof editMemberForm>>({
		image: null,
		email: member.email,
		firstName: member.firstName,
		lastName: member.lastName
	});

	let hasChanged = $derived(
		form.email !== member.email ||
			form.firstName !== member.firstName ||
			form.lastName !== member.lastName ||
			form.image !== null
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
			console.error(e);
			toaster.addToast('Kunne ikke opdatere medlem', 'Noget gik galt', 'error');
		}
	};

	let fullName = $derived(`${member.firstName} ${member.lastName}`);
</script>

<form onsubmit={handleSubmit} class="flex flex-col gap-8">
	<div class="flex flex-col items-center gap-8 md:flex-row">
		<ProfilePictureSelector bind:file={form.image} imageUrl={member.profilePictureUrl} />
		<div class="flex flex-col items-center space-y-4 md:items-start">
			<div class="flex flex-col items-center space-y-1 md:items-start">
				<h1 class="text-2xl font-semibold">{fullName}</h1>
				<span class="text-text/50 text-center md:text-left"
					>{teams.map(({ displayName }) => displayName).join(', ')}</span
				>
			</div>
			<MemberStatusIndicator status={member.active ? 'approved' : 'non-approved'} />
		</div>
	</div>
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
