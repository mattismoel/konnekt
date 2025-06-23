import type { Member } from "../features/auth/member"
import { type Team, type TeamType } from "../features/auth/team"

const includedTeamNames: TeamType[] = [
	"project-leader",
	"public-relations",
	"event-management",
	"booking",
	"visual-identity",
	"economy",
]

type Props = {
	allTeams: Team[]
	members: Member[]
}

const TeamDisplay = ({ members, allTeams }: Props) => {
	const includedTeams = allTeams.filter(team => includedTeamNames.includes(team.name))

	const includedMembers = members
		.filter(member => member.teams.some(team =>
			includedTeams.some(includedTeam => includedTeam.name === team.name)
		))

	const specialMembers: Member[] = []
	const regularMembers: Member[] = []

	includedMembers.forEach(m => m.specialRole ? specialMembers.push(m) : regularMembers.push(m))

	return (
		<div className="@container flex flex-col gap-16">
			<div className="grid grid-cols-1 gap-8 @3xl:grid-cols-2">
				{specialMembers.map(member => <MemberInfo key={member.id} member={member} includedTeams={includedTeams} />)}
				{regularMembers.map(member => <MemberInfo key={member.id} member={member} includedTeams={includedTeams} />)}
			</div>
		</div>
	)
}

type MemberInfoProps = {
	member: Member
	includedTeams: Team[]
}

const MemberInfo = ({ member, includedTeams }: MemberInfoProps) => {
	const memberTeams = member.teams.filter(team => includedTeams.some(t => t.name === team.name))
	const teamsString = memberTeams.map(t => t.displayName).join(", ")

	return (
		<a href={`mailto:${member.email}`} className="group rounded-sm overflow-hidden bg-gradient-to-br from-zinc-700 to-zinc-900 p-[1px] transition-colors hover:from-zinc-600 hover:to-zinc-900">
			<div className="flex flex-col @lg:flex-row bg-gradient-to-tl from-zinc-950 to-zinc-950 rounded-sm overflow-hidden transition-colors group-hover:to-zinc-900">
				<img
					src={member.profilePictureUrl}
					className="w-full aspect-video @lg:w-32 @lg:block object-cover shrink-0"
				/>
				<div className="py-4 px-8">
					<div className="flex flex-col gap-2">
						<span className="font-semibold line-clamp-1">{member.firstName} {member.lastName}</span>
						<div className="text-text/50 flex flex-col">
							<span className="line-clamp-1">
								{member.specialRole ? member.specialRole : teamsString}
							</span>
							<span className="line-clamp-1">{member.email}</span>
						</div>
					</div>
				</div>
			</div>
		</a>
	)
}

export default TeamDisplay
