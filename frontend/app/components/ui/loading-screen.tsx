import { Spinner } from "./spinner"

type Props = {
  label?: string;
}

export const LoadingScreen = ({ label = "Loader..." }: Props) => {
  return (
    <div className="bg-background fixed h-dvh w-screen top-0 left-0 flex justify-center items-center">
      <div className="flex flex-col gap-2 justify-center items-center">
        <Spinner />
        <p>{label}</p>
      </div>
    </div>
  )
}
