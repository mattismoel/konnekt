import type { MetaFunction } from "@remix-run/node";

export const meta: MetaFunction = () => {
  return [
    { title: "Konnekt" },
    { name: "description", content: "Konnekt hjemmeside" },
  ];
};

export default function Index() {
  return (
    <main className="h-sub-nav">
      Welcome to konnekt
    </main>
  );
}
