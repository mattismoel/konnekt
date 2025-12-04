<script lang="ts">
	import { doTest } from "./test.remote";
</script>

<main class="flex min-h-svh flex-col items-center justify-center">
	<form {...doTest} enctype="multipart/form-data" class="flex flex-col">
		<input {...doTest.fields.fileField.as("file")} />
		<button>Submit</button>
		{doTest.fields
			.allIssues()
			?.map((i) => i.message)
			.join(",")}
	</form>
</main>
