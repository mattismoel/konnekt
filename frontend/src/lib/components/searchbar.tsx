import type { InputHTMLAttributes } from "react";
import { FaSearch } from "react-icons/fa";

type Props = Omit<InputHTMLAttributes<HTMLInputElement>, "value" | "onChange"> & {
	search: string;
	onChange: (newSearch: string) => void
};

const Searchbar = ({ search, placeholder = "Søg...", onChange, ...rest }: Props) => (
	<div className="relative w-full">
		<input
			{...rest}
			placeholder={placeholder}
			value={search}
			onChange={e => onChange(e.target.value)}
			type="text"
			className="h-full w-full min-w-48 rounded-sm border-zinc-800 bg-zinc-900 py-2 pr-4 pl-12"
		/>
		<FaSearch className="text-text/75 absolute top-1/2 left-4 -translate-y-1/2" />
	</div>
)

export default Searchbar
