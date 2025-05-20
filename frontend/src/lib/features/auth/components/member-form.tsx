import { Controller, FormProvider, useForm, useFormContext, useFormState } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { editMember, memberForm, type Member, type MemberFormValues } from '../member';
import type { Team } from '../team';
import { useAuth } from '@/lib/context/auth';
import ProfilePictureSelector from '@/lib/components/profile-picture-selector';
import MemberStatusIndicator from '@/lib/components/member-status-indicator';
import FormField from '@/lib/components/form-field';
import Input from '@/lib/components/ui/input';
import Button from '@/lib/components/ui/button/button';
import { createSubmitHandler } from '@/lib/api';
import Picker, { type Entry } from '@/lib/components/ui/picker';
import { createContext, useContext, useState } from 'react';
import PillList from '@/lib/components/pill-list';
import { FaPen } from 'react-icons/fa6';
import { useQueryClient } from '@tanstack/react-query';
import { useNavigate } from '@tanstack/react-router';

type MemberFormContext = {
	member: Member;
	memberTeams: Team[]
	isCurrentMember: boolean;

	teams: Team[]
}

const MemberFormContext = createContext<MemberFormContext | undefined>(undefined)

const useMemberFormContext = () => {
	const ctx = useContext(MemberFormContext)
	if (!ctx) throw new Error("No MemberFormContext.Provider found!")

	return ctx
}

type Props = {
	member: Member;
	memberTeams: Team[]

	teams: Team[];
};

const MemberForm = ({ member, memberTeams, teams }: Props) => {
	const queryClient = useQueryClient()
	const navigate = useNavigate()

	const { member: currentMember } = useAuth()

	let isCurrentMember = currentMember?.id === member.id;

	const methods = useForm({
		defaultValues: {
			image: undefined,
			firstName: member.firstName,
			lastName: member.lastName,
			email: member.email,
			memberTeams: memberTeams.map(({ id }) => id),
		},
		resolver: zodResolver(memberForm),
	})

	const {
		formState: { errors, isDirty },
		setValue,
		handleSubmit,
	} = methods

	let fullName = `${member.firstName} ${member.lastName}`;

	const onSubmit = createSubmitHandler({
		action: async (form: MemberFormValues) => {
			await editMember(member.id, form)
			await queryClient.invalidateQueries({ queryKey: ["members"] })
			navigate({ to: "/admin/members" })
		},
		errorMessage: "Kunne ikke redigere medlem",
		successMessage: "Medlem redigeret",
		navigateTo: "/admin/members",
	})

	return (
		<MemberFormContext.Provider value={{ member, teams, memberTeams, isCurrentMember }}>
			<FormProvider {...methods}>
				<form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-8">
					<div className="w-full flex flex-col items-center gap-8 md:flex-row">
						<FormField error={errors.image} className='justify-center'>
							<ProfilePictureSelector
								onChange={(newFile) => setValue("image", newFile)}
								src={member.profilePictureUrl}
							/>
						</FormField>
						<div className="flex flex-col items-center space-y-4 md:items-start">
							<div className="flex flex-col items-center space-y-1 md:items-start">
								<h1 className="text-2xl font-semibold">{fullName}</h1>
								<span className="text-text/50 text-center md:text-left"
								>{memberTeams.map(({ displayName }) => displayName).join(', ')}</span
								>
							</div>
							<MemberStatusIndicator status={member.active ? 'approved' : 'non-approved'} />
						</div>
					</div>

					<GeneralSection />
					<TeamsSection />

					<Button type="submit" disabled={!isDirty}>Opdatér</Button>
				</form>
			</FormProvider>
		</MemberFormContext.Provider>
	)
}

const GeneralSection = () => {
	const { formState: { errors }, register } = useFormContext<MemberFormValues>()
	const { isCurrentMember } = useMemberFormContext()

	return (
		<section>
			<h1 className="text-2xl font-bold font-heading mb-4">Generelt</h1>

			<div className="flex flex-col gap-4">
				<div className="flex gap-4">
					<FormField error={errors.firstName}>
						<Input {...register("firstName")} placeholder="Fornavn" disabled={!isCurrentMember} />
					</FormField>
					<FormField error={errors.lastName}>
						<Input {...register("lastName")} placeholder="Efternavn" disabled={!isCurrentMember} />
					</FormField>
				</div>

				<FormField error={errors.email}>
					<Input {...register("email")} type="email" placeholder="Email" disabled={!isCurrentMember} />
				</FormField>
			</div>
		</section>
	)
}

const TeamsSection = () => {
	const { control, formState: { errors }, watch } = useFormContext<MemberFormValues>()
	const { teams } = useMemberFormContext()

	const [showPicker, setShowPicker] = useState(false)

	const entries: Entry[] = teams.map(({ id, displayName }) => ({
		id: id.toString(),
		value: id.toString(),
		name: displayName,
	}))

	const selected = entries.filter(entry =>
		watch("memberTeams").includes(parseInt(entry.value))
	)

	return (
		<section>
			<h1 className="font-bold font-heading mb-8 text-2xl">Hold</h1>

			<FormField error={errors.memberTeams}>
				<PillList entries={selected.map(entry => entry.name)}>
					<Button variant="ghost" onClick={() => setShowPicker(true)} className="h-10 rounded-full px-4" ><FaPen />Vælg</Button>
				</PillList>
			</FormField>

			<FormField error={errors.memberTeams}>
				<Controller
					control={control}
					name="memberTeams"
					render={({ field: { onChange } }) => (
						<Picker
							title="Vælg medlemshold..."
							description="Her kan du vælge de medlemshold, som medlemmet associeres med."
							entries={entries}
							selected={selected}
							show={showPicker}
							onClose={() => setShowPicker(false)}
							onChange={(selectedEntries) =>
								onChange(selectedEntries.map(({ value }) => parseInt(value)))
							}
						/>
					)}
				/>
			</FormField>
		</section>
	)
}

export default MemberForm
