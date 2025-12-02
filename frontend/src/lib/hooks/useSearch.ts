import { useEffect, useState } from "react";

type StringKeys<T> = {
  [K in keyof T]: T[K] extends string ? K : never;
}[keyof T];

export const useSearch = <T>(entries: T[], searchProperty: StringKeys<T>) => {
  const [search, setSearch] = useState("");

  const [results, setResults] = useState(entries);

  useEffect(() => {
    if (search === "") {
      setResults(entries);
      return;
    }

    setResults(
      entries.filter((entry) => {
        const value = entry[searchProperty] as string;
        return value.toLowerCase().includes(search.toLowerCase());
      }),
    );
  }, [search]);

  return { search, setSearch, results };
};
