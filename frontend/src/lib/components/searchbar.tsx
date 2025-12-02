import type { InputHTMLAttributes } from "react";
import { FaSearch } from "react-icons/fa";
import { cn } from "../clsx";

type Props = Omit<
  InputHTMLAttributes<HTMLInputElement>,
  "value" | "onChange"
> & {
  search: string;
  onChange: (newSearch: string) => void;
};

const Searchbar = ({
  search,
  placeholder = "Søg...",
  onChange,
  ...rest
}: Props) => (
  <div className={cn("relative w-full", rest.className)}>
    <input
      {...rest}
      placeholder={placeholder}
      value={search}
      onChange={(e) => onChange(e.target.value)}
      type="text"
      className="h-full w-full min-w-48 rounded-full border-zinc-800 bg-zinc-900 py-2 pr-6 pl-14"
    />
    <FaSearch className="text-text/75 absolute top-1/2 left-6 -translate-y-1/2" />
  </div>
);

export default Searchbar;
