<script lang="ts">
	import Card from "$lib/components/Card.svelte";
	import ErrorList from "$lib/components/ErrorList.svelte";
	import FormField from "$lib/components/FormField.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import type { PageProps } from "./$types";

	let { form }: PageProps = $props();
</script>

<main class="flex min-h-svh items-center justify-center">
	<form method="POST" enctype="multipart/form-data">
		<Card
			title="Tilmeld"
			description="Her kan du tilmelde dig som medlem af foreningen Konnekt."
			class="max-w-lg"
		>
			<div class="flex flex-col gap-6">
				<FormField issues={form?.fieldErrors.avatar}>
					<Input name="avatar" type="file" accept="image/*" />
				</FormField>

				<div class="flex flex-col gap-2">
					<div class="flex gap-4">
						<FormField issues={form?.fieldErrors.firstName}>
							<Input
								name="firstName"
								placeholder="Fornavn"
								class="w-full"
								value={form?.data.firstName}
							/>
						</FormField>
						<FormField issues={form?.fieldErrors.lastName}>
							<Input
								name="lastName"
								placeholder="Efternavn"
								class="w-full"
								value={form?.data.lastName}
							/>
						</FormField>
					</div>
					<FormField issues={form?.fieldErrors.email}>
						<Input type="email" name="email" placeholder="Email" value={form?.data.email} />
					</FormField>
				</div>

				<div class="flex flex-col gap-2">
					<FormField issues={form?.fieldErrors.password}>
						<Input type="password" name="password" placeholder="Adgangskode" />
					</FormField>
					<FormField issues={form?.fieldErrors.passwordConfirm}>
						<Input type="password" name="passwordConfirm" placeholder="Gentag adgangskode" />
					</FormField>
				</div>
				{#if form?.formErrors}
					<ErrorList issues={form?.formErrors} />
				{/if}

				<Button type="submit" class="w-full">Registrér</Button>
			</div>
		</Card>
	</form>
</main>
