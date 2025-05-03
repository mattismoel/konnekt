<script lang="ts">
	import AvatarImage from '$lib/assets/avatar.png';
	import { page } from '$app/state';

	import { cn } from '$lib/clsx';

	import { type Member } from '$lib/features/auth/member';

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
	import { hasSomeTeam } from '$lib/features/auth/team';

	type Props = {
		member: Member;
		expanded: boolean;
		onToggle: () => void;
	};

	let { member, expanded, onToggle }: Props = $props();

	// TODO: Implement sign out functionality.
	const handleSignOut = () => {
		return;
	};
</script>

<aside
	class:expanded
	class={cn(
		'group w-sidenav-lg fixed z-50 flex h-full -translate-x-full flex-col gap-y-8 border-r border-zinc-900 bg-zinc-950 p-8 transition-transform duration-100 [.expanded]:translate-x-0'
	)}
>
	<!-- LOGO SECTION -->
	<div class="flex items-center justify-between">
		<Logo class="h-5" />
		<Button onclick={onToggle} variant="ghost" class="h-12 w-12">
			<CollapseIcon class="rotate-180 group-[.expanded]:rotate-0" />
		</Button>
	</div>

	<ul class="flex flex-1 flex-col gap-2">
		{#if hasSomeTeam(member.teams, ['admin', 'event-management'])}
			{@render entry(EventIcon, '/admin/events', 'Events')}
		{/if}
		{#if hasSomeTeam(member.teams, ['admin', 'booking'])}
			{@render entry(ArtistIcon, '/admin/artists', 'Kunstnere')}
		{/if}
		{#if hasSomeTeam(member.teams, ['admin', 'event-management'])}
			{@render entry(VenueIcon, '/admin/venues', 'Venues')}
		{/if}
		{#if hasSomeTeam(member.teams, ['admin', 'team-management'])}
			{@render entry(MemberIcon, '/admin/members', 'Medlemmer')}
		{/if}
		{@render entry(SettingsIcon, '/admin/general', 'Generelt')}
	</ul>
	{@render memberInformation(member)}
</aside>

{#snippet memberInformation(member: Member)}
	{@const teamsString = member.teams.map((t) => t.displayName).join(', ')}

	<div class="flex items-center justify-between gap-8">
		<a href="/admin/members/{member.id}" class="flex items-center gap-6">
			<img
				src={member.profilePictureUrl || AvatarImage}
				alt="Profile"
				class="h-10 w-10 rounded-full object-cover"
			/>
			<div>
				<span class="line-clamp-1">{member.firstName} {member.lastName}</span>
				<span class="text-text/50 line-clamp-1" title={teamsString}>{teamsString}</span>
			</div>
		</a>
		<button type="button" title="Log ud" onclick={handleSignOut}>
			<SignOutIcon class="text-text/50 hover:text-text text-xl" />
		</button>
	</div>
{/snippet}

{#snippet entry(Logo: Component, href: string, label: string)}
	<li
		class:is-active={page.url.pathname === href}
		class="[.is-active]:text-text text-text/50 hover:text-text/75 flex w-full rounded-sm border border-transparent transition-colors duration-75 hover:border-zinc-900 [.is-active]:border-zinc-800 [.is-active]:bg-zinc-900"
	>
		<a {href} class="flex w-full items-center gap-2 px-4 py-2" title={label}>
			<Logo />
			<span>{label}</span>
		</a>
	</li>
{/snippet}
