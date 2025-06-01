import { Outlet, createRootRouteWithContext } from '@tanstack/react-router'
import { QueryClient } from "@tanstack/react-query"

type RouteProps = {
  queryClient: QueryClient
}

export const Route = createRootRouteWithContext<RouteProps>()({
  component: Outlet,
})
