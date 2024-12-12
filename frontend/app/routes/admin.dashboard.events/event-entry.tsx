import { format } from "date-fns";
import { BiTrash } from "react-icons/bi";
import { EventDTO } from "@/lib/dto/event.dto";
import { formatDateString } from "@/lib/time";

type Props = {
  event: EventDTO;
  onDelete: () => void;
}

export const EventEntry = ({ event, onDelete }: Props) => {
  const { id, title, fromDate } = event;
  const formattedDate = `${formatDateString(fromDate)}, ${format(fromDate, "kk:mm")}`

  return (
    <div className="px-4 py-2 flex justify-between hover:bg-neutral-900">
      <a href={`/admin/event/rediger/${id}`} className="w-full flex">
        <span className="flex-1">{title}</span>
        <span className="hidden sm:block flex-1">{formattedDate}</span>
      </a>
      <button className="text-lg hover:text-red-500" onClick={onDelete}><BiTrash /></button>
    </div>
  )
}
