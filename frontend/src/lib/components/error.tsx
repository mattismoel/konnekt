import { useRouterState, type ErrorComponentProps } from "@tanstack/react-router"
import { APIError } from "../api"
import type { PropsWithChildren } from "react"
import NotFoundComponent from "./not-found"


const ErrorComponent = ({ error }: ErrorComponentProps) => {
	return (
		<main className="min-h-svh flex flex-col justify-center items-center">
			{(error instanceof APIError)
				? error.status === 404
					? <NotFoundComponent />
					: <Display status={error.status}>{error.cause}</Display>
				: <Display status={500}>{error.message}</Display>
			}
		</main >
	)
}

type DisplayProps = {
	status: number;
}

const Display = ({ status, children }: PropsWithChildren<DisplayProps>) => {
	return (
		<div className="flex flex-col">
			<span className="underline">{status}</span>
			<p className="text-text/75">
				{children}
			</p>
		</div>
	)
}


export default ErrorComponent
