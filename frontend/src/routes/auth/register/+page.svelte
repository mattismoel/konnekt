<script lang="ts">
	import Card from "$lib/components/Card.svelte";
	import ErrorList from "$lib/components/ErrorList.svelte";
	import FormField from "$lib/components/FormField.svelte";
	import Button from "$lib/components/ui/Button.svelte";
	import Input from "$lib/components/ui/Input.svelte";
	import { register } from "$lib/features/auth/auth.remote";
</script>

<main class="flex min-h-svh items-center justify-center">
	<form {...register} enctype="multipart/form-data">
		<Card
			title="Tilmeld"
			description="Her kan du tilmelde dig som medlem af foreningen Konnekt."
			class="max-w-lg"
		>
			<div class="flex flex-col gap-6">
				<FormField issues={register.fields.avatar.issues()?.map((i) => i.message)}>
					<Input {...register.fields.avatar.as("file")} />
				</FormField>

				<div class="flex flex-col gap-2">
					<div class="flex gap-4">
						<FormField issues={register.fields.firstName.issues()?.map((i) => i.message)}>
							<Input
								{...register.fields.firstName.as("text")}
								class="w-full"
								placeholder="Fornavn"
							/>
						</FormField>
						<FormField issues={register.fields.lastName.issues()?.map((i) => i.message)}>
							<Input
								{...register.fields.lastName.as("text")}
								class="w-full"
								placeholder="Efternavn"
							/>
						</FormField>
					</div>
					<FormField issues={register.fields.email.issues()?.map((i) => i.message)}>
						<Input {...register.fields.email.as("email")} placeholder="Email" />
					</FormField>
				</div>

				<div class="flex flex-col gap-2">
					<FormField issues={register.fields.password.issues()?.map((i) => i.message)}>
						<Input {...register.fields.password.as("password")} placeholder="Adgangskode" />
					</FormField>
					<FormField issues={register.fields.passwordConfirm.issues()?.map((i) => i.message)}>
						<Input
							{...register.fields.passwordConfirm.as("password")}
							placeholder="Gentag adgangskode"
						/>
					</FormField>
				</div>

				{#if register.fields.issues()}
					<ErrorList issues={register.fields.issues()?.map((i) => i.message)} />
				{/if}

				<Button type="submit" class="w-full">Registrér</Button>
			</div>
		</Card>
	</form>
</main>
