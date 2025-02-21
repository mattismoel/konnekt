<script lang="ts">
	import { page } from '$app/state';
	import { cn } from '$lib/clsx';
	import type { User } from '$lib/user';
	import SettingsIcon from '~icons/mdi/settings-outline';

	let { children, data } = $props();
	let { user, roles } = $derived(data);
</script>

<main class="h-sub-nav px-auto grid w-screen grid-cols-[290px_1fr] gap-4 bg-zinc-950 py-20">
	{@render sidebar()}
	{@render children()}
</main>

{#snippet sidebar()}
	<div class="flex flex-col rounded-sm border border-zinc-800 bg-zinc-900 p-2">
		<div class="flex flex-1 flex-col p-4">
			<span class="mb-4 text-xl font-black">KONNEKT&reg;</span>
			<ul class="flex-1 space-y-1">
				{@render sidebarEntry('Events.', '/admin/dashboard/events')}
				{@render sidebarEntry('Kunstenre.', '/admin/dashboard/artists')}
				{@render sidebarEntry('Generelt.', '/admin/dashboard/general')}
			</ul>
		</div>
		{@render userInformation(user)}
	</div>
{/snippet}

{#snippet sidebarEntry(title: string, url: string)}
	<li
		class={cn('text-zinc-500', {
			'font-bold text-zinc-100': page.url.pathname === url
		})}
	>
		<a href={url}>{title}</a>
	</li>
{/snippet}

{#snippet userInformation(user: User)}
	<div
		class="group flex items-center gap-4 rounded-md border border-transparent p-2 transition-colors hover:border-zinc-700 hover:bg-zinc-800"
	>
		<img
			src="https://upload.wikimedia.org/wikipedia/commons/8/89/Portrait_Placeholder.png"
			alt=""
			class="h-10 w-10 rounded-full"
		/>
		<div class="flex flex-col text-sm">
			<span class="line-clamp-1">{user.firstName} {user.lastName}</span>
			<p class="line-clamp-1 text-zinc-500">{roles.map((r) => r.displayName).join(', ')}</p>
		</div>
		<a href={`/admin/dashboard/user/${user.id}`}>
			<SettingsIcon class="text-zinc-500 hover:text-zinc-300" />
		</a>
	</div>
{/snippet}
