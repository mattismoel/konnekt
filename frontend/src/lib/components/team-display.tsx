import type { Member } from "../features/auth/member";
import { type Team, type TeamType } from "../features/auth/team";

const includedTeamNames: TeamType[] = [
  "project-leader",
  "public-relations",
  "event-management",
  "booking",
  "visual-identity",
  "economy",
];

type Props = {
  allTeams: Team[];
  members: Member[];
};

const TeamDisplay = ({ members, allTeams }: Props) => {
  const includedTeams = allTeams.filter((team) =>
    includedTeamNames.includes(team.name),
  );

  const includedMembers = members.filter((member) =>
    member.teams.some((team) =>
      includedTeams.some((includedTeam) => includedTeam.name === team.name),
    ),
  );

  const specialMembers: Member[] = [];
  const regularMembers: Member[] = [];

  includedMembers.forEach((m) =>
    m.specialRole ? specialMembers.push(m) : regularMembers.push(m),
  );

  return (
    <div className="@container flex flex-col">
      <div className="mb-16 grid grid-cols-1 gap-2 @3xl:grid-cols-2">
        {specialMembers.map((member) => (
          <MemberInfo
            key={member.id}
            member={member}
            includedTeams={includedTeams}
          />
        ))}
      </div>

      <h4 className="mb-12 text-center text-text/50">
        Øvrige {regularMembers.length > 5 && regularMembers.length} medlemmer
      </h4>

      <div className="flex flex-wrap justify-center gap-x-20 gap-y-12">
        {regularMembers.map((member) => (
          <div className="flex flex-col items-center">
            <p className="mb-1 text-base font-medium">
              {member.firstName} {member.lastName}
            </p>
            <p className="text-xs text-text/50">
              {member.teams
                .slice(0, 2)
                .map((team) => team.displayName)
                .join(", ")}
            </p>
          </div>
        ))}
      </div>
    </div>
  );
};

type MemberInfoProps = {
  member: Member;
  includedTeams: Team[];
};

const MemberInfo = ({ member, includedTeams }: MemberInfoProps) => {
  const memberTeams = member.teams.filter((team) =>
    includedTeams.some((t) => t.name === team.name),
  );
  const teamsString = memberTeams.map((t) => t.displayName).join(", ");

  return (
    <div className="group flex flex-col overflow-hidden rounded-sm border border-zinc-800 bg-zinc-900 transition-colors @lg:flex-row">
      <img
        src={member.profilePictureUrl}
        className="aspect-video w-full shrink-0 object-cover @lg:block @lg:w-32"
      />
      <div className="flex max-w-full flex-col gap-2 overflow-hidden px-8 py-4">
        <span className="line-clamp-1 font-medium">
          {member.firstName} {member.lastName}
        </span>
        <div className="w-full text-text/50">
          <span className="line-clamp-1">
            {member.specialRole ? member.specialRole : teamsString}
          </span>
        </div>
      </div>
    </div>
  );
};

export default TeamDisplay;
