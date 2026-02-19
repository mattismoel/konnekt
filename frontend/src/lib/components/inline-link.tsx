import type { AnchorHTMLAttributes } from "react";
import { cn } from "../clsx";

type Props = AnchorHTMLAttributes<HTMLAnchorElement>;

const InlineLink = ({ children, ...rest }: Props) => (
  <a
    {...rest}
    className={cn(
      "font-medium underline transition-colors duration-100 hover:text-text-light hover:underline",
      rest.className,
    )}
  >
    {children}
  </a>
);

export default InlineLink;
