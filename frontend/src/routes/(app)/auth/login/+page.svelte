<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';
	import { authStore } from '$lib/auth.svelte';

	import Button from '$lib/components/ui/Button.svelte';
	import * as Card from '$lib/components/ui/card/index';
	import Input from '$lib/components/ui/Input.svelte';
	import { login, loginForm } from '$lib/features/auth/auth';
	import { memberPermissions } from '$lib/features/auth/permission';
	import { memberTeams } from '$lib/features/auth/team';
	import type { z } from 'zod';

	let form = $state<z.infer<typeof loginForm>>({
		email: '',
		password: ''
	});

	const handleLogin = async () => {
		try {
			const member = await login(fetch, form);
			const teams = await memberTeams(fetch, member.id);
			const permissions = await memberPermissions(fetch, member.id);

			authStore.auth = { member, permissions, teams };
		} catch (e) {
			throw e;
		}

		goto('/admin/events');
	};
</script>

<main class="flex h-svh items-center justify-center">
	<form onsubmit={handleLogin}>
		<Card.Root class="max-w-96">
			<Card.Header>
				<Card.Title>Login</Card.Title>
				<Card.Description>Her kan du logge ind som medlem på Konnekts dashboard.</Card.Description>
			</Card.Header>

			<Card.Content>
				<section class="flex flex-col gap-4">
					<Input type="email" bind:value={form.email} placeholder="Email" />
					<Input type="password" bind:value={form.password} placeholder="Adgangskode" />
				</section>
			</Card.Content>

			<Card.Footer>
				<Button class="w-full" type="submit">Login</Button>
			</Card.Footer>
		</Card.Root>
	</form>
</main>
