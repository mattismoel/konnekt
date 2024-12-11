import { useMemo } from "react";
import { GenreSelectorEntry } from "./genre-selector-entry";

type Props = {
  genres: string[];
  selected: string[];
  onChange: (newGenres: string[]) => void
}

export const GenreSelector = ({ genres, selected, onChange }: Props) => {
  const toggleSelected = (genre: string, isSelected: boolean) => {
    const updatedSelected = isSelected ? selected.filter(g => g !== genre) : [...selected, genre]

    console.log(updatedSelected)
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
