<script lang="ts">
	import ProfilePictureSelector from '$lib/components/ProfilePictureSelector.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import * as Card from '$lib/components/ui/card/index';
	import Input from '$lib/components/ui/Input.svelte';
	import { register, type registerForm } from '$lib/features/auth/auth';
	import type { z } from 'zod';

	const form = $state<z.infer<typeof registerForm>>({
		email: '',
		firstName: '',
		lastName: '',
		password: '',
		passwordConfirm: '',
		profilePictureUrl: ''
	});

	let profilePictureFile = $state<File | null>(null);

	const submit = async (e: SubmitEvent) => {
		e.preventDefault();

		await register(fetch, form, profilePictureFile);
	};
</script>

<form onsubmit={submit}>
	<Card.Root class="max-w-lg">
		<Card.Header>
			<Card.Title>Tilmeld</Card.Title>
			<Card.Description>Her kan du tilmelde dig som medlem af foreningen Konnekt.</Card.Description>
		</Card.Header>

		<Card.Content class="gap-16">
			<div class="flex w-full justify-center">
				<ProfilePictureSelector bind:file={profilePictureFile} />
			</div>

			<div class="flex flex-col gap-4">
				<div class="flex gap-4">
					<Input placeholder="Fornavn" bind:value={form.firstName} />
					<Input placeholder="Efternavn" bind:value={form.lastName} />
				</div>
				<Input type="email" placeholder="Email" bind:value={form.email} />
			</div>

			<div class="flex flex-col gap-4">
				<Input type="password" placeholder="Adgangskode" bind:value={form.password} />
				<Input type="password" placeholder="Gentag adgangskode" bind:value={form.passwordConfirm} />
			</div>
		</Card.Content>

		<Card.Footer>
			<Button type="submit" class="w-full">Registrér</Button>
		</Card.Footer>
	</Card.Root>
</form>
