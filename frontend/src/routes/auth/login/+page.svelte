<script lang="ts">
	import { goto } from '$app/navigation';
	import { PUBLIC_BACKEND_URL } from '$env/static/public';

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

<main class="min-h-sub-nav">
	<form class="flex flex-col gap-2" onsubmit={handleLogin}>
		<input type="email" bind:value={email} placeholder="Email" />
		<input type="password" bind:value={password} placeholder="Adgangskode" />
		<button>Login</button>
	</form>
</main>
