import type { PropsWithChildren } from 'react';
import Searchbar from './searchbar';
import List from './list/list';

type Props = PropsWithChildren<{
	search: string;
	onChange: (newSearch: string) => void
}>;

const SearchList = ({ search, onChange, children }: Props) => (
	<div className="flex flex-col gap-8">
		<Searchbar search={search} onChange={onChange} />
		<List>
			{children}
		</List>
	</div>
)

export default SearchList
