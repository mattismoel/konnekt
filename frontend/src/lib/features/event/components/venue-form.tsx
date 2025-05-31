import { zodResolver } from '@hookform/resolvers/zod';

import { Controller, useForm } from 'react-hook-form';
import { useQueryClient } from '@tanstack/react-query';

import { COUNTRIES_MAP } from '../countries';
import { createSubmitHandler } from '@/lib/api';
import { createVenue, editVenue, venueForm, type Venue, type VenueFormValues } from '../venue';

import Button from '@/lib/components/ui/button/button';
import FormField from '@/lib/components/form-field';
import Input from '@/lib/components/ui/input';
import Selector from '@/lib/components/ui/selector';
import { useAuth } from '@/lib/context/auth';

type Props = {
	venue?: Venue
}

const VenueForm = ({ venue }: Props) => {
	const { hasPermissions } = useAuth()
	const queryClient = useQueryClient()

	const isEditable = hasPermissions(["edit:venue"])

	const {
		control,
		formState: { errors },
		register,
		handleSubmit,
	} = useForm({
		disabled: !isEditable,
		defaultValues: { ...venue },
		resolver: zodResolver(venueForm),
	})

	const onSubmit = createSubmitHandler({
		errorMessage: venue ? "Kunne ikke redigere venue" : "Kunne ikke skabe venue",
		successMessage: venue ? "Venue redigeret" : "Venue skabt",
		navigateTo: "/admin/venues",
		action: async (form: VenueFormValues) => {
			venue ? await editVenue(venue.id, form) : await createVenue(form)
			await queryClient.invalidateQueries({ queryKey: ["venues"] })
		}
	})

	return (
		<form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-8">
			<div className="flex flex-col gap-4">
				<FormField error={errors.name}>
					<Input {...register("name")} placeholder="Venuenavn" />
				</FormField>

				<div className="flex gap-4">
					<FormField error={errors.city}>
						<Input {...register("city")} placeholder="By" className="flex-1" />
					</FormField>
					<div>
						<Controller
							control={control}
							name="countryCode"
							render={({ field: { onChange, ...rest } }) => (
								<FormField error={errors.countryCode}>
									<Selector
										{...rest}
										onChange={(e) => onChange(e.target.value)}
										placeholder="Vælg land..."
										className="h-min w-min"
									>
										{Array.from(COUNTRIES_MAP).map(([value, name]) => (
											<option value={value} key={value}>{name}</option>
										))}
									</Selector>
								</FormField>
							)}
						/>
					</div>
				</div>
			</div>
			{isEditable && <Button type="submit" className="w-full">Offentliggør</Button>}
		</form>
	)
}

export default VenueForm
