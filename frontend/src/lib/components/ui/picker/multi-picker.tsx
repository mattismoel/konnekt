import type { Entry as EntryType } from "./entry";
import Entry from "./entry";

type Props = {
  entries: EntryType[];
  selected: EntryType[];
  max?: number;
  onChange: (selected: EntryType[]) => void;
};

const MultiPicker = ({ entries, selected, max, onChange }: Props) => {
  const handleToggle = (entry: EntryType) => {
    const isSelected = selected.some((e) => e.id === entry.id);

    if (max && !isSelected && selected.length >= max) return;

    const newSelected = isSelected
      ? selected.filter((e) => e.id !== entry.id)
      : [...selected, entry];

    onChange(newSelected);
  };

  return (
    <div className="flex flex-col gap-1">
      {entries.map((entry) => (
        <Entry
          key={entry.id}
          selected={selected.some((e) => e.id === entry.id)}
          onToggle={() => handleToggle(entry)}
        >
          {entry.name}
        </Entry>
      ))}
    </div>
  );
};

export default MultiPicker;
