import { useEffect, useState } from "react"

type ScrollPosition = {
  x: number;
  y: number;
}

export const useScroll = () => {
  const [scrollPos, setScrollPos] = useState<ScrollPosition>({ x: 0, y: 0 })

  useEffect(() => {
    const handleScroll = () => {
      setScrollPos({
        x: window.pageXOffset,
        y: window.pageYOffset,
      })
    }

    window.addEventListener("scroll", handleScroll)

    return () => {
      window.removeEventListener("scroll", handleScroll)
    }
  }, [])

  return scrollPos
}
