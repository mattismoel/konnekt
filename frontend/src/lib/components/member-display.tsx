import type { Member } from "@/lib/features/auth/member"

type Props = {
	members: Member[];
}

const TeamDisplay = ({ members }: Props) => {
	return (
		<div>
			{members.map(({ firstName, lastName, email, avatarUrl, teams, specialRole }) => (
				<div className="@container group isolate">
					<div className="flex flex-col @xl:flex-row relative border border-zinc-900 overflow-hidden rounded-sm transition-colors duration-600 group-hover:border-zinc-800">
						<img src={avatarUrl} className="z-10 max-h-80 @xl:h-32 @xl:aspect-square @xl:w-auto w-full object-cover group-hover:brightness-105 transition-[filter]" />
						<img src={avatarUrl} className="-z-50 absolute top-0 left-0 h-full w-full object-cover opacity-0 blur-2xl scale-150 transition-opacity duration-300 group-hover:opacity-5" />

						<div className="p-4 flex flex-col z-20">
							<span className="text-heading font-semibold mb-2">{firstName} {lastName}</span>
							<span>{specialRole ?? teams.map(({ displayName }) => displayName).join(", ")}</span>
							<a href={`mailto:${email}`} className="hover:underline hover:text-heading">{email}</a>
						</div>
					</div>
				</div>
			))}
		</div>
	)
}

export default TeamDisplay;
