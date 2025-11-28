import type { LiHTMLAttributes, PropsWithChildren } from "react";

type Props = {
  entries: string[];
};

const PillList = ({ entries, children }: PropsWithChildren<Props>) => (
  <ul className="flex flex-wrap items-center gap-2">
    {children}
    {entries.map((name) => (
      <Pill key={name}>{name}</Pill>
    ))}
  </ul>
);

const Pill = ({ children, ...rest }: LiHTMLAttributes<HTMLLIElement>) => (
  <li
    {...rest}
    className="flex h-10 w-fit cursor-default items-center justify-center rounded-full border border-zinc-800 bg-zinc-900 px-4 text-text/75"
  >
    {children}
  </li>
);

export default PillList;
