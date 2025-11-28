import { Link } from "@tanstack/react-router";

const NotFoundComponent = () => (
  <main className="mx-responsive flex h-svh w-screen items-center justify-center bg-background py-16">
    <div className="flex flex-col items-center">
      <h1 className="mb-4 text-center font-heading text-7xl font-bold">404</h1>
      <span className="mb-4">Hov! Denne side findes desværre ikke...</span>
      <Link
        to="/"
        className="text-text/75 underline decoration-text/50 decoration-wavy decoration-2 underline-offset-4 transition-colors duration-300 hover:text-text hover:decoration-text"
      >
        Lad os følge dig hjem igen...
      </Link>
    </div>
  </main>
);

export default NotFoundComponent;
