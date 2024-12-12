import { useMemo, useState } from "react";
import { GenreSelectorEntry } from "./genre-selector-entry";

type Props = {
  genres: string[];
  defaultSelected: string[];
  onChange: (newGenres: string[]) => void
}

export const GenreSelector = ({ genres, defaultSelected, onChange }: Props) => {
  const [selected, setSelected] = useState<string[]>(defaultSelected)

  const toggleSelected = (genre: string, isSelected: boolean) => {
    const updatedSelected = isSelected ? selected.filter(g => g !== genre) : [...selected, genre]

    setSelected(updatedSelected)
    onChange(updatedSelected)
  }

  return (
    <div className="flex flex-wrap gap-3">
      {genres.map(genre => {
        const isSelected = useMemo(() => selected.includes(genre), [selected])
        return (
          <GenreSelectorEntry
            key={genre}
            genre={genre}
            onToggle={() => toggleSelected(genre, isSelected)}
            isSelected={isSelected}
          />
        )
      })}
    </div>
  )
}
