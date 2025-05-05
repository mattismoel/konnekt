<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	import Button from '$lib/components/ui/Button.svelte';
	import * as Card from '$lib/components/ui/card/index';
	import Input from '$lib/components/ui/Input.svelte';

	let email = $state('');
	let password = $state('');

	const handleLogin = async () => {
		const res = await fetch(`${PUBLIC_BACKEND_URL}/auth/login`, {
			credentials: 'include',
			body: JSON.stringify({ email, password }),
			method: 'post'
		});

		if (!res.ok) {
			throw new Error('Could not login');
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
					<Input type="email" bind:value={email} placeholder="Email" />
					<Input type="password" bind:value={password} placeholder="Adgangskode" />
				</section>
			</Card.Content>

			<Card.Footer>
				<Button class="w-full" type="submit">Login</Button>
			</Card.Footer>
		</Card.Root>
	</form>
</main>
