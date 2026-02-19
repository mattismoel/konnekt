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

      <h4 className="text-text/50 mb-12 text-center">
        Øvrige {regularMembers.length > 5 && regularMembers.length} medlemmer
      </h4>

      <div className="flex flex-wrap justify-center gap-x-20 gap-y-12">
        {regularMembers.map((member) => (
          <RegularMemberInfo
            key={member.id}
            member={member}
            includedTeams={includedTeams}
          />
        ))}
      </div>
    </div>
  );
};

type MemberInfoProps = { member: Member; includedTeams: Team[] };

const MemberInfo = ({ member, includedTeams }: MemberInfoProps) => {
  const memberTeams = member.teams.filter((team) =>
    includedTeams.some((t) => t.name === team.name),
  );
  const teamsString = memberTeams.map((t) => t.displayName).join(", ");

  return (
    <div className="group flex items-center overflow-hidden rounded-full border border-zinc-800 bg-zinc-900 p-1 transition-colors">
      <img
        src={member.profilePictureUrl}
        className="aspect-square w-24 shrink-0 rounded-full object-cover @lg:w-32"
      />
      <div className="h-mi flex w-full flex-col gap-2 overflow-hidden px-8 py-4">
        <span className="line-clamp-1 font-medium text-heading @lg:text-xl">
          {member.firstName} {member.lastName}
        </span>
        <div className="w-full">
          <span className="line-clamp-1">
            {member.specialRole ? member.specialRole : teamsString}
          </span>
        </div>
      </div>
    </div>
  );
};

type RegularMemberInfoProps = {
  member: Member;
  includedTeams: Team[];
};

const RegularMemberInfo = ({
  member,
  includedTeams,
}: RegularMemberInfoProps) => {
  const memberTeams = member.teams.filter((team) =>
    includedTeams.some((t) => t.name === team.name),
  );
  const teamsString = memberTeams.map((t) => t.displayName).join(", ");

  return (
    <div className="flex flex-col items-center">
      <p className="mb-1 text-base font-medium">
        {member.firstName} {member.lastName}
      </p>
      <p className="text-text/50 text-xs">{teamsString}</p>
    </div>
  );
};

export default TeamDisplay;
