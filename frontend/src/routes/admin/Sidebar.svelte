<script lang="ts">
	import { page } from '$app/state';

	import { cn } from '$lib/clsx';

	import { hasSomeRole, type Role } from '$lib/auth';
	import { type User } from '$lib/auth';

	import Logo from '$lib/assets/Logo.svelte';

	import EventIcon from '~icons/mdi/event';
	import VenueIcon from '~icons/mdi/warehouse';
	import ArtistIcon from '~icons/mdi/people-group';
	import MemberIcon from '~icons/mdi/people';
	import SettingsIcon from '~icons/mdi/settings';

	import SignOutIcon from '~icons/mdi/sign-out';

	import CollapseIcon from '~icons/mdi/arrow-collapse-left';
	import { type Component } from 'svelte';
	import Button from '$lib/components/ui/Button.svelte';

	type Props = {
		user: User;
		roles: Role[];
		expanded: boolean;
		onToggle: () => void;
	};

	let { user, roles, expanded, onToggle }: Props = $props();
</script>

<aside
	class:expanded
	class={cn(
		'group w-sidenav-sm [.expanded]:w-sidenav-lg group fixed z-50 flex h-full flex-col items-center gap-y-8 border-r border-zinc-900 p-8 px-0 [.expanded]:items-stretch [.expanded]:px-8'
	)}
>
	<div class="flex items-baseline justify-between">
		<Logo class="hidden h-5 group-[.expanded]:block" />
		<Button onclick={onToggle} variant="ghost" class="px-2"
			><CollapseIcon class="rotate-180 group-[.expanded]:rotate-0" /></Button
		>
	</div>
	<section class="flex flex-1 flex-col gap-8">
		<ul class="space-y-1">
			{#if hasSomeRole(roles, ['admin', 'event-management'])}
				{@render entry(EventIcon, '/admin/dashboard/events', 'Events')}
			{/if}
			{#if hasSomeRole(roles, ['admin', 'event-management'])}
				{@render entry(VenueIcon, '/admin/dashboard/venues', 'Venues')}
			{/if}
			{#if hasSomeRole(roles, ['admin', 'booking'])}
				{@render entry(ArtistIcon, '/admin/dashboard/artists', 'Kunstnere')}
			{/if}
			{#if hasSomeRole(roles, ['admin', 'team-management'])}
				{@render entry(MemberIcon, '/admin/dashboard/members', 'Medlemmer')}
			{/if}
			{@render entry(SettingsIcon, '/admin/dashboard/general', 'Generelt')}
		</ul>
	</section>
	{@render userInformation(user, roles)}
</aside>

{#snippet userInformation(user: User, roles: Role[])}
	{@const rolesString = roles.map((r) => r.displayName).join(', ')}
	<div class="flex flex-col-reverse items-center justify-between gap-8 group-[.expanded]:flex-row">
		<a href="/admin/dashboard/user/{user.id}" class="flex items-center gap-4">
			<img
				src="https://upload.wikimedia.org/wikipedia/commons/8/89/Portrait_Placeholder.png"
				alt=""
				class="aspect-square h-full w-10 rounded-full"
			/>
			<div class="hidden group-[.expanded]:block">
				<span>{user.firstName} {user.lastName}</span>
				<span class="text-text/50 line-clamp-1" title={rolesString}>{rolesString}</span>
			</div>
		</a>
		<button type="button" title="Log ud">
			<SignOutIcon class="text-text/50 hover:text-text text-xl" />
		</button>
	</div>
{/snippet}

{#snippet entry(Logo: Component, href: string, label: string)}
	<li
		class={cn(
			'text-text/75 hover:text-text flex aspect-square w-full items-center justify-center rounded-md border border-transparent group-[.expanded]:aspect-auto hover:border-zinc-800',
			{
				'text-text border-zinc-800 bg-zinc-900': page.url.pathname === href
			}
		)}
	>
		<a {href} class="flex w-full items-center gap-2 px-4 py-2" title={label}>
			<Logo />
			<span class="hidden group-[.expanded]:inline">{label}</span>
		</a>
	</li>
{/snippet}
