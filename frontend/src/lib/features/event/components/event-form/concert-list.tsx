import { FaArrowRight, FaArrowsRotate, FaPlus, FaXmark } from "react-icons/fa6";
import { useEventFormContext } from "./event-form";
import { Controller } from "react-hook-form";
import type { PropsWithChildren } from "react";
import Button from "@/lib/components/ui/button/button";
import Card from "@/lib/components/ui/card";
import Selector from "@/lib/components/ui/selector";
import LinkButton from "@/lib/components/ui/button/link-button";
import DatetimePicker from "@/lib/components/datetime-picker";

const ConcertList = ({ children }: PropsWithChildren) => {
	const { onAddConcert } = useEventFormContext()

	return (
		<div className="flex flex-col gap-4">
			{children}
			<Button variant="ghost" onClick={onAddConcert}><FaPlus />Tilføj</Button>
		</div>
	)
}

type EntryProps = { index: number }

const Entry = ({ index }: EntryProps) => {
	const { artists, fields, control, setValue, onDeleteConcert } = useEventFormContext()

	return (
		<Card className="relative">
			<Card.Header>
				<Card.Title>#{index + 1}</Card.Title>
				<button type="button" onClick={() => onDeleteConcert(index)} className="absolute top-4 right-4"><FaXmark /></button>
			</Card.Header>
			<Card.Content className="gap-8 @container">
				<div className="flex gap-4">
					<Selector
						onChange={(e) => setValue(`concerts.${index}.artistID`, parseInt(e.target.value))}
						defaultValue={artists.find(artist => artist.id === fields[index].artistID)?.id} placeholder="Vælg kunstner..." className="w-full">
						{artists.map(artist => (
							<option value={artist.id} key={artist.id}>{artist.name}</option>
						))}
					</Selector>
					<div className="flex gap-2">
						<Button variant="ghost" className="aspect-square h-full"><FaArrowsRotate /></Button>
						<LinkButton to="/admin/artists/create" className="aspect-square h-full"><FaPlus /></LinkButton>
					</div>
				</div>
				<div className="flex flex-col gap-4 @xl:flex-row @xl:gap-8 items-center">
					<Controller
						control={control}
						name={`concerts.${index}.from`}
						render={({ field: { onChange, value } }) => (
							<DatetimePicker value={value} onChange={onChange} />
						)}
					/>
					<FaArrowRight className="hidden shrink-0 @xl:block" />
					<Controller
						control={control}
						name={`concerts.${index}.to`}
						render={({ field: { onChange, value } }) => (
							<DatetimePicker value={value} onChange={onChange} />
						)}
					/>
				</div>
			</Card.Content>
		</Card>
	)
}

ConcertList.Entry = Entry

export default ConcertList
