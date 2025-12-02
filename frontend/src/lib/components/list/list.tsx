import {
  forwardRef,
  type AnchorHTMLAttributes,
  type PropsWithChildren,
} from "react";
import { createLink } from "@tanstack/react-router";
import { cn } from "@/lib/clsx";

type Props = AnchorHTMLAttributes<HTMLAnchorElement> & {
  disabled?: boolean;
};

const List = ({ children }: PropsWithChildren) => (
  <ul className="flex flex-col gap-2">{children}</ul>
);

const Entry = createLink(
  forwardRef<HTMLAnchorElement, Props>(
    ({ disabled = false, className, ...rest }, ref) => (
      <li>
        <a
          ref={ref}
          {...rest}
          className={cn(
            "flex items-center justify-between rounded-xl border border-zinc-800 bg-zinc-900",
          )}
        />
      </li>
    ),
  ),
);

List.Entry = Entry;

export default List;
