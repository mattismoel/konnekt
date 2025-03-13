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
	<Card>
		<h1 class="mb-8 text-2xl font-bold">Login.</h1>
		<form class="flex flex-col gap-2" onsubmit={handleLogin}>
			<Input type="email" bind:value={email} label="Email" />
			<Input type="password" bind:value={password} label="Adgangskode" />
			<Button expandX type="submit">Login</Button>
		</form>
	</Card>
</main>
