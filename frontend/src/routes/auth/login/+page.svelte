<script lang="ts">
	import Card from "$lib/components/Card.svelte";
	import FormField from "$lib/components/FormField.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import { login } from "$lib/features/auth/auth.remote";
</script>

<main class="flex min-h-svh items-center justify-center">
	<Card
		title="Log ind"
		description="Her kan du logge ind som medlem på Konnekts dashboard."
		class="max-w-sm"
	>
		<form {...login} class="flex flex-col">
			<div class="mb-8 flex flex-col gap-2">
				<FormField>
					<Input {...login.fields.email.as("email")} placeholder="Email" />
				</FormField>
				<FormField>
					<Input {...login.fields.password.as("password")} placeholder="Adganskode" />
				</FormField>
			</div>
			<Button>Log ind</Button>
		</form>
	</Card>
</main>
