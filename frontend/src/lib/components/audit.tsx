import { Link } from "@tanstack/react-router";
import type { Member } from "../features/auth/member";
import { formatRelative } from "date-fns";
import { useAuth } from "../context/auth";

type Props = {
  updatedByMember: Member;
  updatedAt: Date;
};

const formatName = (m: Member): string => `${m.firstName} ${m.lastName}`;

const Audit = ({ updatedByMember, updatedAt }: Props) => {
  const { member } = useAuth();
  return (
    <div className="flex items-center gap-4">
      <img
        src={updatedByMember.profilePictureUrl}
        className="aspect-square h-6 rounded-full"
      />
      <span className="text-text/75">
        Redigeret af&nbsp;
        {member?.id === updatedByMember.id ? (
          "dig"
        ) : (
          <Link
            to="/admin/members/$memberId"
            params={{ memberId: updatedByMember.id.toString() }}
            className="hover:underline"
          >
            {formatName(updatedByMember)}
          </Link>
        )}
        &nbsp;
        {formatRelative(updatedAt, new Date())}
      </span>
    </div>
  );
};

export default Audit;
