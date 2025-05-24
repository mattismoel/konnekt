import LandingImagesForm from '@/lib/components/landing-images-form'
import { landingImagesQueryOptions } from '@/lib/features/content/query'
import { useSuspenseQuery } from '@tanstack/react-query'
import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/admin/content/')({
  component: RouteComponent,
  loader: async ({ context: { queryClient } }) => {
    queryClient.ensureQueryData(landingImagesQueryOptions)
  }
})

function RouteComponent() {
  const { data: images } = useSuspenseQuery(landingImagesQueryOptions)

  return (
    <LandingImagesForm images={images} />
  )
}
