import Input from './ui/input';
import { useEffect, useState } from 'react';
import type { WithClassName } from '../type';
import { cn } from '../clsx';

type Props = WithClassName<{
	value?: Date;
	onChange?: (newDate: Date) => void;
}>;

const DatetimePicker = ({ value, className, onChange, ...rest }: Props) => {
	const [date, setDate] = useState(() => value
		? value.toISOString().slice(0, 10)
		: undefined
	)

	const [time, setTime] = useState(() => value ? value.toTimeString().slice(0, 5) : undefined)

	useEffect(() => {
		if (!date || !time) return

		const [hours, minutes] = time.split(":").map(part => parseInt(part))

		let combined = new Date(date)

		combined.setHours(hours)
		combined.setMinutes(minutes)
		combined.setSeconds(0)
		combined.setMilliseconds(0)

		onChange?.(combined)
	}, [date, time])

	return (
		<div className={cn("w-full flex gap-2", className)}>
			<Input
				{...rest}
				className="w-full"
				type="date"
				value={date}
				onChange={(e) => setDate(e.target.value)}
			/>
			<Input
				{...rest}
				className="w-min"
				type="time"
				value={time}
				onChange={(e) => setTime(e.target.value)}
			/>
		</div>
	)
}

export default DatetimePicker
