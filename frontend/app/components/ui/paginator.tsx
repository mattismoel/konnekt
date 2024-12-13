import { cn } from "@/lib/utils";

type Props = {
  totalPages: number;
  currentPage: number;
  onSelect: (page: number) => void;
}

export const Paginator = ({ totalPages, currentPage, onSelect }: Props) => {
  return (
    <div className="w-full flex justify-center">
      {[...Array(totalPages)].map((_, i) => {
        const page = i + 1
        return (
          <button
            className={cn("px-4 py-1", { "font-bold bg-zinc-900 rounded-md": page === currentPage })}
            onClick={() => onSelect(page)}
          >
            {page}
          </button>
        )
      })}
    </div>
  )
}
