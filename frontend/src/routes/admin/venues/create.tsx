import { createFileRoute } from "@tanstack/react-router";

import VenueForm from "@/lib/features/event/components/venue-form";

export const Route = createFileRoute("/admin/venues/create")({
  component: RouteComponent,
});

function RouteComponent() {
  return (
    <main className="mx-responsive flex min-h-svh items-center justify-center py-32">
      <div>
        <h1 className="mb-4 font-heading text-2xl font-bold">Lav venue</h1>
        <VenueForm />
      </div>
    </main>
  );
}
