<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

	import Button from '$lib/components/ui/Button.svelte';
	import Card from '$lib/components/ui/Card.svelte';
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

		goto('/admin/dashboard/events');
	};
</script>

<main class="flex h-svh items-center justify-center">
	<Card class="max-w-96">
		<form class="flex w-full flex-col gap-8" onsubmit={handleLogin}>
			<section>
				<h1 class="font-heading mb-4 text-2xl font-bold">Login</h1>
				<p class="text-text/50">Her kan du logge ind som medlem på Konnekts dashboard.</p>
			</section>
			<section>
				<Input type="email" bind:value={email} label="Email" />
				<Input type="password" bind:value={password} label="Adgangskode" />
			</section>
			<Button class="w-full" type="submit">Login</Button>
		</form>
	</Card>
</main>
