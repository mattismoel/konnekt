import { cn } from "~/lib/utils";

type Props = {
  genre: string;
  isSelected: boolean;
  onToggle: () => void;
}

export const GenreSelectorEntry = ({ genre, isSelected, onToggle }: Props) => {
  return (
    <button
      type="button"
      className={cn(
        "px-2 py-1 rounded-sm border transition-colors text-foreground/20",
        { "bg-zinc-800 text-foreground": isSelected }
      )}
      onClick={() => onToggle()}
    >
      {genre}
    </button>
  )
}
