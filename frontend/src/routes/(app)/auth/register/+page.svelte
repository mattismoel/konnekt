<script lang="ts">
	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
	import Input from '$lib/components/ui/Input.svelte';
	import { register, type registerForm } from '$lib/features/auth/auth';
	import type { z } from 'zod';

	const form = $state<z.infer<typeof registerForm>>({
		email: '',
		firstName: '',
		lastName: '',
		password: '',
		passwordConfirm: ''
	});

	const submit = async (e: SubmitEvent) => {
		e.preventDefault();

		await register(fetch, form);
	};
</script>

<main class="flex h-svh items-center justify-center">
	<form onsubmit={submit}>
		<Card class="max-w-lg space-y-8">
			<div>
				<h1 class="mb-4 text-2xl font-bold">Tilmeld</h1>
				<span class="text-text/50">
					Her kan du tilmelde dig som medlem af foreningen Konnekt.
				</span>
			</div>
			<section>
				<div class="flex gap-2">
					<Input label="Fornavn" bind:value={form.firstName} />
					<Input label="Efternavn" bind:value={form.lastName} />
				</div>
				<Input type="email" label="Email" bind:value={form.email} />
				<Input type="password" label="Adgangskode" bind:value={form.password} />
				<Input type="password" label="Gentag adgangskode" bind:value={form.passwordConfirm} />
			</section>
			<Button type="submit" class="w-full">Registrér</Button>
		</Card>
	</form>
</main>
