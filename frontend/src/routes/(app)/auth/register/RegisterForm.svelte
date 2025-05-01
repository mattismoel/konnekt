<script lang="ts">
	import ProfilePictureSelector from '$lib/components/ProfilePictureSelector.svelte';
	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import { register, type registerForm } from '$lib/features/auth/auth';
	import { uploadMemberProfilePicture } from '$lib/features/auth/member';
	import type { z } from 'zod';

	const form = $state<z.infer<typeof registerForm>>({
		email: '',
		firstName: '',
		lastName: '',
		password: '',
		passwordConfirm: '',
		profilePictureUrl: "",
	});

	let profilePictureFile = $state<File | null>(null);

	$inspect(profilePictureFile)

	const submit = async (e: SubmitEvent) => {
		e.preventDefault();

		await register(fetch, form, profilePictureFile);
	};
</script>

<form onsubmit={submit}>
	<Card class="flex max-w-lg flex-col gap-16">
		<div>
			<h1 class="mb-4 text-2xl font-bold">Tilmeld</h1>
			<span class="text-text/50"> Her kan du tilmelde dig som medlem af foreningen Konnekt. </span>
		</div>
		<section class="flex justify-center">
			<ProfilePictureSelector bind:file={profilePictureFile} />
		</section>
		<section class="flex flex-col gap-8">
			<div class="flex gap-2">
				<Input placeholder="Fornavn" bind:value={form.firstName} />
				<Input placeholder="Efternavn" bind:value={form.lastName} />
			</div>

			<Input type="email" placeholder="Email" bind:value={form.email} />

			<div class="flex flex-col gap-4">
				<Input type="password" placeholder="Adgangskode" bind:value={form.password} />
				<Input type="password" placeholder="Gentag adgangskode" bind:value={form.passwordConfirm} />
			</div>
		</section>
		<Button type="submit" class="w-full">Registrér</Button>
	</Card>
</form>
