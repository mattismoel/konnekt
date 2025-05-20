import { zodResolver } from '@hookform/resolvers/zod';

import { useForm } from 'react-hook-form';
import { useQueryClient } from '@tanstack/react-query';

import { COUNTRIES_MAP } from '../countries';
import { createSubmitHandler } from '@/lib/api';
import { createVenue, editVenue, venueForm, type Venue, type VenueFormValues } from '../venue';

import Button from '@/lib/components/ui/button/button';
import FormField from '@/lib/components/form-field';
import Input from '@/lib/components/ui/input';
import Selector from '@/lib/components/ui/selector';

type Props = {
	venue?: Venue
	disabled?: boolean;
}

const VenueForm = ({ venue, disabled = false }: Props) => {
	const queryClient = useQueryClient()

	const {
		formState: { errors },
		register,
		handleSubmit,
		getValues,
		setValue,
	} = useForm({
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
					<Input {...register("name")} disabled={disabled} placeholder="Venuenavn" />
				</FormField>

				<div className="flex gap-4">
					<FormField error={errors.city}>
						<Input {...register("city")} disabled={disabled} placeholder="By" className="flex-1" />
					</FormField>
					<div>
						<FormField error={errors.countryCode}>
							<Selector
								onChange={(e) => setValue("countryCode", e.target.value)}
								disabled={disabled}
								className="h-min w-min"
								placeholder="Vælg land..."
								defaultValue={getValues("countryCode")}
							>
								{Array.from(COUNTRIES_MAP).map(([value, name]) => (
									<option value={value} key={value}>{name}</option>
								))}
							</Selector>
						</FormField>
					</div>
				</div>
			</div>
			{!disabled && <Button type="submit" className="w-full">Offentliggør</Button>}
		</form>
	)
}

export default VenueForm
