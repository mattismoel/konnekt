import { addHours, differenceInHours, differenceInMilliseconds, format, set, setMinutes } from "date-fns";
import { capitalise } from "./string";

export const formatDateString = (date: Date): string => {
  let weekday = date.toLocaleDateString("da", { weekday: "long" });

  weekday = capitalise(weekday)

  return `${weekday} / ${format(date, "dd.MM.yy")}`;
}

export const getFullHoursSurroudningDates = (earlier: Date, later: Date): number => {
  earlier = set(earlier, { year: 1970, month: 1, date: 1 })
  later = set(later, { year: 1970, month: 1, date: 1 })

  earlier = setMinutes(earlier, 0)

  if (later.getMinutes() > 0) {
    later = addHours(later, 1)
  }

  later = setMinutes(later, 0)

  const diffHours = differenceInHours(later, earlier)

  return diffHours
}

export const formatHoursAsTimestamp = (hours: number): string => {
  const paddedHours = hours.toString().padStart(2, "0");
  return `${paddedHours}:00`;
};

/**
  * Returns the difference between two times of day in milliseconds. This means
  * that only the two dates time of day will be taken into consideration.
  *
  * @example - {04/05/2013 10:00} and {06/02/2025 12:00} -> 2 hours -> 7.2e6 milliseconds
  */
export const distanceBetweenTimesOfDay = (d1: Date, d2: Date): number => {
  const t1 = set(d1, { year: 1970, month: 1, date: 1 });
  const t2 = set(d2, { year: 1970, month: 1, date: 1 });

  const diff = differenceInMilliseconds(t2, t1)

  return diff
}

export const sleep = async (ms: number = 1000): Promise<void> => {
  return new Promise(resolve => setTimeout(resolve, ms))
}
