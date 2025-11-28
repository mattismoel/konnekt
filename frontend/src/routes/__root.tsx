import { Outlet, createRootRouteWithContext } from "@tanstack/react-router";
import { QueryClient } from "@tanstack/react-query";
import { da } from "date-fns/locale";
import { setDefaultOptions } from "date-fns";
import { useEffect } from "react";

type RouteProps = {
  queryClient: QueryClient;
};

export const Route = createRootRouteWithContext<RouteProps>()({
  component: () => {
    useEffect(() => setDefaultOptions({ locale: da }), []);
    return <Outlet />;
  },
});
