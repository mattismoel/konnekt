import { cn } from "../clsx";

type Props = {
  entries: {
    id: string;
    value: string;
    ignorePrefix?: boolean;
  }[];
  selected: string;
  prefix?: string;

  className?: string;

  onSwitch: (newId: string) => void;
};

const Switch = ({ entries, selected, prefix, onSwitch, ...rest }: Props) => {
  return (
    <div
      className={cn(
        "@container flex w-full items-center justify-between gap-8",
        rest.className,
      )}
    >
      <div
        className={cn(
          "flex w-full flex-col items-center justify-between gap-4 overflow-hidden rounded-md border border-transparent",
          "@2xl:flex-row @2xl:rounded-full @2xl:border-zinc-900 @2xl:p-1",
          prefix && "w-full @2xl:pl-8",
        )}
      >
        <span
          className={cn(
            "w-full text-center text-sm whitespace-nowrap text-text/75 italic",
            "@2xl:text-left",
          )}
        >
          {prefix}
        </span>

        <div className="flex w-full flex-col gap-2 @2xl:flex-row @2xl:gap-1">
          {entries.map((entry) => (
            <button
              key={entry.id}
              type="button"
              onClick={() => onSwitch(entry.id)}
              title={entry.value}
              className={cn(
                "rounded-full border border-zinc-900 px-6 py-2 whitespace-nowrap transition-colors duration-100 hover:border-zinc-800 hover:bg-zinc-900",
                "@2xl:border-transparent @2xl:text-sm",
                "before:invisible before:block before:h-0 before:overflow-hidden before:font-semibold before:content-[attr(title)]",
                selected === entry.id &&
                  "bg-zinc-100 font-semibold text-zinc-900 hover:border-zinc-100 hover:bg-zinc-100",
              )}
            >
              {entry.value}
            </button>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Switch;
