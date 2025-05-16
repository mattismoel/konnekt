import { useEffect, useState } from "react"
import { useScroll } from "./useScroll";

type MousePosition = { x: number; y: number };

export const useMousePos = () => {
  const scrollPos = useScroll()
  const [rawMousePos, setRawMousePos] = useState<MousePosition>({ x: 0, y: 0 })

  const mousePos = {
    x: rawMousePos.x + scrollPos.x,
    y: rawMousePos.y + scrollPos.y,
  }

  useEffect(() => {
    setRawMousePos(prev => ({
      x: prev.x + scrollPos.x,
      y: prev.y + scrollPos.y
    }))
  }, [scrollPos])

  useEffect(() => {
    const handleMouseMove = (e: MouseEvent) => {
      setRawMousePos({
        x: e.clientX,
        y: e.clientY
      })
    }

    window.addEventListener("mousemove", handleMouseMove)

    return () => {
      window.removeEventListener("mousemove", handleMouseMove)
    }
  }, [])

  return mousePos
}
