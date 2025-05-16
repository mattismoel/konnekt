import { useEffect, type RefObject } from "react";

type EventType =
  | 'mousedown'
  | 'mouseup'
  | 'touchstart'
  | 'touchend'
  | 'focusin'
  | 'focusout'

export const useOnClickOutside = <T extends HTMLElement = HTMLElement>(
  ref: RefObject<T> | RefObject<T>[] | RefObject<null>,
  handler: (e: MouseEvent | TouchEvent | FocusEvent) => void,
  eventType: EventType = "mousedown",
  eventListenerOptions: AddEventListenerOptions = {}
): void => {
  useEffect(() => {
    if (!ref) return

    const handleEvent = (e: MouseEvent | TouchEvent | FocusEvent) => {
      const target = e.target as Node
      if (!target || !target.isConnected) return

      if (Array.isArray(ref)) {
        const isOutside = ref
          .filter(r => Boolean(r.current))
          .every(r => r.current && !r.current.contains(target))

        if (isOutside) handler(e)
        return
      }

      const isOutside = ref.current && !ref.current.contains(target)
      if (isOutside) handler(e)
    }

    window.addEventListener(eventType, handleEvent, eventListenerOptions)

    return () => {
      window.removeEventListener(eventType, handleEvent, eventListenerOptions)
    }
  })
}
