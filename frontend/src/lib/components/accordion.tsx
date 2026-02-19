import { useState, type PropsWithChildren } from "react";
import { cn } from "../clsx";
import { FaChevronDown } from "react-icons/fa6";

type Props = {
  title: string;
};

const Accordion = ({ title, children }: PropsWithChildren<Props>) => {
  const [expanded, setExpanded] = useState(false);

  return (
    <div
      className={cn(
        "group overflow-hidden rounded-xl border border-zinc-800 hover:text-text-light",
        expanded && "expanded text-text-light",
      )}
    >
      <button
        type="button"
        className="flex w-full items-center gap-6 bg-zinc-900 px-8 py-3 text-left text-base transition-colors group-[.expanded]:bg-zinc-800 hover:bg-zinc-800"
        onClick={() => setExpanded((prev) => !prev)}
      >
        <FaChevronDown className="shrink-0 text-sm transition-transform group-[.expanded]:rotate-180" />
        <span className="font-medium">{title}</span>
      </button>

      <div className="hidden cursor-default overflow-hidden px-8 py-8 group-[.expanded]:block">
        <p className="prose leading-loose text-text-light-muted prose-invert prose-a:text-text-light-muted">
          {children}
        </p>
      </div>
    </div>
  );
};

export default Accordion;
