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
			{#if hasSomeTeam(member.teams, ['admin', 'event-management'])}
				{@render entry(EventIcon, '/admin/dashboard/events', 'Events')}
			{/if}
			{#if hasSomeTeam(member.teams, ['admin', 'booking'])}
				{@render entry(ArtistIcon, '/admin/dashboard/artists', 'Kunstnere')}
			{/if}
			{#if hasSomeTeam(member.teams, ['admin', 'event-management'])}
				{@render entry(VenueIcon, '/admin/dashboard/venues', 'Venues')}
			{/if}
			{#if hasSomeTeam(member.teams, ['admin', 'team-management'])}
				{@render entry(MemberIcon, '/admin/dashboard/members', 'Medlemmer')}
			{/if}
			{@render entry(SettingsIcon, '/admin/dashboard/general', 'Generelt')}
		</ul>
	</section>
	{@render memberInformation(member)}
</aside>

{#snippet memberInformation(member: Member)}
	{@const teamsString = member.teams.map((t) => t.displayName).join(', ')}

	<div class="flex flex-col-reverse items-center justify-between gap-8 group-[.expanded]:flex-row">
		<a href="/admin/dashboard/member/{member.id}" class="flex items-center gap-4">
			<img
				src={member.profilePictureUrl || AvatarImage}
				alt=""
				class="aspect-square h-full w-10 rounded-full object-cover"
			/>
			<div class="hidden group-[.expanded]:block">
				<span class="line-clamp-1">{member.firstName} {member.lastName}</span>
				<span class="text-text/50 line-clamp-1" title={teamsString}>{teamsString}</span>
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
