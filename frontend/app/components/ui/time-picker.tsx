import { ChangeEvent, useState } from "react";
import { Input } from "./input";
import { cn } from "~/lib/utils";

type Time = {
  hours: number;
  minutes: number;
}


type Props = {
  time: Time
  onChange: (time: Time) => void;
  className?: string
}

export const TimePicker = ({ time, onChange, className }: Props) => {
  const onHoursChange = (e: ChangeEvent<HTMLInputElement>) => {
    let newValue = parseInt(e.currentTarget.value)

    newValue = Math.min(newValue, 23)
    newValue = Math.max(newValue, 0)

    onChange({ ...time, hours: newValue })
    padZero(e, newValue)
  }

  const onMinutesChange = (e: ChangeEvent<HTMLInputElement>) => {
    let newValue = parseInt(e.currentTarget.value)

    newValue = Math.min(newValue, 59)
    newValue = Math.max(newValue, 0)

    onChange({ ...time, minutes: newValue })
    padZero(e, newValue)
  }

  return (
    <div className={cn("flex justify-center items-center gap-2", className)}>
      <Input
        className="flex-1 text-center max-w-16"
        defaultValue={time.hours.toString().padStart(2, "0")}
        min={0}
        max={24}
        onChange={onHoursChange}
      />
      <span>:</span>
      <Input
        className="flex-1 text-center max-w-16"
        defaultValue={time.minutes.toString().padStart(2, "0")}
        onChange={onMinutesChange}
      />
    </div>
  )
}

const padZero = (e: ChangeEvent<HTMLInputElement>, value: number) => {
  const maxLength = parseInt(e.currentTarget.getAttribute("maxLength") || "2")
  const isNegative = value < 0

  let newValue = (
    "0".repeat(maxLength)
    + Math.abs(value).toString())
    .slice(-maxLength)

  if (isNegative) {
    newValue = "-" + newValue
  }

  e.currentTarget.value = newValue
}
