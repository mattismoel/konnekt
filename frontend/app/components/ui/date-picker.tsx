import { format, setHours, setMinutes } from "date-fns"
import { Calendar as CalendarIcon } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Calendar } from "@/components/ui/calendar"
import { Popover, PopoverContent, PopoverTrigger } from "@/components/ui/popover"
import { useState } from "react"
import { TimePicker } from "./time-picker"

type Props = {
  initialDate?: Date;
  className?: string;
  placeholder?: string;
}

type Time = { hours: number, minutes: number }

export function DateTimePicker({ initialDate, placeholder = "Pick a date", className }: Props) {
  const [date, setDate] = useState(initialDate)
  const [time, setTime] = useState<Time>({
    hours: date?.getHours() || new Date().getHours(),
    minutes: date?.getMinutes() || new Date().getMinutes()
  })

  const changeTime = (newTime: Time) => {
    const withMinutes = setMinutes(date || new Date(), newTime.minutes)
    const withHours = setHours(withMinutes, newTime.hours)

    setTime(newTime)
    setDate(withHours)
  }

  const changeDate = (newDate: Date | undefined) => {
    const withMinutes = setMinutes(newDate || new Date(), time.minutes)
    const withHours = setHours(withMinutes, time.hours)

    setDate(withHours)
  }

  return (
    <Popover>
      <PopoverTrigger asChild>
        <Button
          variant={"outline"}
          className={cn(
            "w-[280px] justify-start text-left font-normal",
            className,
            !date && "text-muted-foreground"
          )}
        >
          <CalendarIcon className="mr-2 h-4 w-4" />
          {date ? format(date, "PPP, HH:mm") : <span>{placeholder}</span>}
        </Button>
      </PopoverTrigger>
      <PopoverContent className="w-auto p-0 divide-y">
        <Calendar
          mode="single"
          selected={date}
          onSelect={changeDate}
          initialFocus
        />
        <TimePicker
          className="w-full p-3"
          time={time}
          onChange={changeTime}
        />
      </PopoverContent>
    </Popover>
  )
}
