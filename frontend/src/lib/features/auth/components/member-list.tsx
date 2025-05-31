import List from "@/lib/components/list/list";
import Avatar from "@/lib/assets/avatar.png"

import { approveMember, deleteMember, type Member } from "../member";
import { useToast } from "@/lib/context/toast";
import { useQueryClient } from "@tanstack/react-query";
import { APIError } from "@/lib/api";
import { FaCheckDouble, FaTrash } from "react-icons/fa";
import MemberStatusIndicator from "@/lib/components/member-status-indicator";
import { useState } from "react";
import ContextMenu from "@/lib/components/context-menu";
import { useAuth } from "@/lib/context/auth";

type Props = {
	members: Member[];
	pendingMembers: Member[];
};

const MemberList = ({ members, pendingMembers }: Props) => {
	const { hasPermissions } = useAuth()

	return (
		<div className="space-y-8">
			{hasPermissions(["edit:member"]) && pendingMembers.length > 0 && (
				<section>
					<h1 className="mb-4">Anmodninger ({pendingMembers.length})</h1>
					<List>
						{pendingMembers.map(member => <ApprovalEntry member={member} />)}
					</List>
				</section>
			)}

			<section>
				<h1 className="mb-4">Medlemmer</h1>
				<List className="space-y-2">
					{members.map(member => <MemberEntry member={member} />)}
				</List>
			</section>
		</div>
	)
}


type MemberEntryProps = {
	member: Member;
};

const MemberEntry = ({ member }: MemberEntryProps) => {
	let [showContextMenu, setShowContextMenu] = useState(false);
	const { addToast } = useToast()
	const { hasPermissions } = useAuth()
	const queryClient = useQueryClient()

	let fullName = `${member.firstName} ${member.lastName}`;

	const handleDeleteMember = async () => {
		if (
			!confirm(
				`Er du sikker på at du vil slette ${fullName} fra foreningen?\n\nHandlingen kan ikke fortrydes.`
			)
		) {
			return;
		}
		try {
			await deleteMember(member.id);
			addToast('Medlem slettet');
			await queryClient.invalidateQueries({ queryKey: ["members"] });
		} catch (e) {
			if (e instanceof APIError) {
				addToast('Kunne ikke slette medlemmet', e.cause, 'error');
				throw e
			}
			addToast('Kunne ikke slette medlemmet', 'Noget gik galt...', 'error');
			throw e
		}
	};

	return (
		<List.Entry className="relative">
			<List.Entry.LinkSection to="/admin/members/$memberId" params={{ memberId: member.id.toString() }} className="flex-row items-center gap-4">
				<img
					src={member.profilePictureUrl || Avatar}
					alt="Profil"
					className="h-8 w-8 rounded-full object-cover"
				/>
				<span className="line-clamp-1">{member.firstName} {member.lastName}</span>
			</List.Entry.LinkSection>

			<List.Entry.Section className="flex-row items-center gap-4 w-min">
				<MemberStatusIndicator
					status={member.active ? 'approved' : 'non-approved'}
					className="hidden md:flex"
				/>
				<ContextMenu.Button onClick={() => setShowContextMenu(true)} />
			</List.Entry.Section>

			<ContextMenu show={showContextMenu} onClose={() => setShowContextMenu(false)}>
				<ContextMenu.LinkEntry
					to="/admin/members/$memberId"
					params={{ memberId: member.id.toString() }}
					disabled={!hasPermissions(['edit:member'])}
				>
					Redigér
				</ContextMenu.LinkEntry>
				<ContextMenu.Entry
					onClick={handleDeleteMember}
					disabled={!hasPermissions(['delete:member'])}
				>
					Slet
				</ContextMenu.Entry>
			</ContextMenu>
		</List.Entry>
	)
}


type ApprovalEntryProps = {
	member: Member;
};

const ApprovalEntry = ({ member }: ApprovalEntryProps) => {
	const { addToast } = useToast()
	const queryClient = useQueryClient()

	const approve = async () => {
		try {
			await approveMember(member.id);
			addToast('Bruger godkendt');
			await queryClient.invalidateQueries({ queryKey: ["members"] })
		} catch (e) {
			if (e instanceof APIError) {
				addToast("Kunne ikke godkende bruger", e.cause, "error")
				throw e
			}

			addToast('Kunne ikke godkende bruger', 'Noget gik galt...', 'error');
			throw e
		}
	};

	const disapprove = async () => {
		try {
			await deleteMember(member.id);
			addToast('Bruger forkastet');
			queryClient.invalidateQueries({ queryKey: ["members"] })
		} catch (e) {
			if (e instanceof APIError) {
				addToast('Kunne ikke forkaste bruger', e.cause, 'error');
				throw e
			}

			addToast('Kunne ikke forkaste bruger', 'Noget gik galt...', 'error');
			throw e
		}
	};

	return (
		<List.Entry className="gap-4">
			<List.Entry.LinkSection to="/admin/members/$memberId" params={{ memberId: member.id.toString() }}>
				<div className="flex flex-1 items-center gap-4">
					<img
						src={member.profilePictureUrl || Avatar}
						alt="Profil"
						className="h-8 w-8 rounded-full object-cover"
					/>
					<span className="line-clamp-1">{member.firstName} {member.lastName} </span>
				</div>
			</List.Entry.LinkSection>

			<List.Entry.Section className="w-min flex-row gap-6">
				<MemberStatusIndicator status="non-approved" className="hidden md:block" />
				<div className="text-text/75 flex gap-2">
					<button className="p-1 hover:text-green-500" onClick={approve}><FaCheckDouble /></button>
					<button className="p-1 hover:text-red-500" onClick={disapprove}><FaTrash /></button>
				</div>
			</List.Entry.Section>
		</List.Entry>
	)
}


export default MemberList
