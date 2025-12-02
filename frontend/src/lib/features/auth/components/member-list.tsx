import List from "@/lib/components/list/list";
import Avatar from "@/lib/assets/avatar.png";

import { approveMember, deleteMember, type Member } from "../member";
import { useToast } from "@/lib/context/toast";
import { useQueryClient } from "@tanstack/react-query";
import { APIError } from "@/lib/api";
import { FaCheckDouble, FaTrash } from "react-icons/fa";
import { useAuth } from "@/lib/context/auth";
import { Link } from "@tanstack/react-router";
import Button from "@/lib/components/ui/button/button";

type Props = {
  members: Member[];
  pendingMembers: Member[];
};

const MemberList = ({ members, pendingMembers }: Props) => {
  const { hasPermissions } = useAuth();

  return (
    <div className="space-y-8">
      {hasPermissions(["edit:member"]) && pendingMembers.length > 0 && (
        <section>
          <h1 className="mb-4">Anmodninger ({pendingMembers.length})</h1>
          <List>
            {pendingMembers.map((member) => (
              <ApprovalEntry member={member} />
            ))}
          </List>
        </section>
      )}

      <section>
        <h1 className="mb-4">Medlemmer</h1>
        <List>
          {members.map((member) => (
            <MemberEntry member={member} />
          ))}
        </List>
      </section>
    </div>
  );
};

type MemberEntryProps = {
  member: Member;
};

const MemberEntry = ({ member }: MemberEntryProps) => {
  return (
    <li className="group relative rounded-full border border-zinc-800 bg-zinc-900 p-1 transition-colors hover:border-zinc-700 hover:bg-zinc-800">
      <Link
        to="/admin/members/$memberId"
        params={{ memberId: member.id.toString() }}
        className="flex w-full items-center gap-4"
      >
        <img
          src={member.profilePictureUrl || Avatar}
          alt="Profil"
          className="aspect-square h-12 rounded-full object-cover"
        />
        <p className="line-clamp-1 font-medium transition-colors group-hover:text-text-light">
          {member.firstName} {member.lastName}
        </p>
      </Link>
    </li>
  );
};

type ApprovalEntryProps = {
  member: Member;
};

const ApprovalEntry = ({ member }: ApprovalEntryProps) => {
  const { addToast } = useToast();
  const queryClient = useQueryClient();

  const approve = async () => {
    try {
      await approveMember(member.id);
      addToast("Bruger godkendt");
      await queryClient.invalidateQueries({ queryKey: ["members"] });
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke godkende bruger", e.cause, "error");
        throw e;
      }

      addToast("Kunne ikke godkende bruger", "Noget gik galt...", "error");
      throw e;
    }
  };

  const disapprove = async () => {
    try {
      await deleteMember(member.id);
      addToast("Bruger forkastet");
      queryClient.invalidateQueries({ queryKey: ["members"] });
    } catch (e) {
      if (e instanceof APIError) {
        addToast("Kunne ikke forkaste bruger", e.cause, "error");
        throw e;
      }

      addToast("Kunne ikke forkaste bruger", "Noget gik galt...", "error");
      throw e;
    }
  };

  return (
    <li className="flex gap-4">
      <Link
        to="/admin/members/$memberId"
        params={{ memberId: member.id.toString() }}
        className="w-full"
      >
        <div className="flex flex-1 items-center gap-4">
          <img
            src={member.profilePictureUrl || Avatar}
            alt="Profil"
            className="h-8 w-8 rounded-full object-cover"
          />
          <span className="line-clamp-1">
            {member.firstName} {member.lastName}{" "}
          </span>
        </div>
      </Link>

      <div className="flex w-min gap-2">
        <Button variant="primary" onClick={approve}>
          <FaCheckDouble /> Godkend
        </Button>
        <Button variant="dangerous" onClick={disapprove}>
          <FaTrash /> Afvis
        </Button>
      </div>
    </li>
  );
};

export default MemberList;
