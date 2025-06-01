import { Link } from "@tanstack/react-router"

const NotFoundComponent = () => (
	<main className="bg-background px-auto flex h-svh w-screen items-center 
		justify-center py-16"
	>
		<div className="flex flex-col items-center">
			<h1 className="mb-4 text-center text-7xl font-heading font-bold">404</h1>
			<span className="mb-4">Hov! Denne side findes desværre ikke...</span>
			<Link to="/" className="text-text/75 decoration-text/50 
				hover:decoration-text hover:text-text underline decoration-wavy
				decoration-2 underline-offset-4 transition-colors duration-300"
			>
				Lad os følge dig hjem igen...
			</Link>
		</div>
	</main>
)

export default NotFoundComponent
