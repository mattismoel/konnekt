import { cn } from "@/lib/utils"

type Direction = "to-bottom" | "to-top" | "to-left" | "to-right"

type Props = {
  direction: Direction
}

export const Fader = ({ direction }: Props) => {
  return (
    <div className={cn(
      "absolute from-transparent to-black",
      { "h-2/3 w-full": direction === "to-top" || direction === "to-bottom" },
      { "h-full w-2/3": direction === "to-left" || direction === "to-right" },

      /* TO TOP*/
      { "bg-gradient-to-t top-0 left-0": direction === "to-top" },

      /* TO BOTTOM*/
      { "bg-gradient-to-b bottom-0 left-0": direction === "to-bottom" },

      /* TO LEFT*/
      { "bg-gradient-to-l top-0 left-0": direction === "to-left" },

      /* TO RIGHT*/
      { "bg-gradient-to-r top-0 right-0": direction === "to-right" },
    )}></div>
  )
}
