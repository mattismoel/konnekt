import { cn } from "@/lib/clsx";
import { Link, useLocation } from "@tanstack/react-router";
import { useScroll } from "@/lib/hooks/useScroll";
import type { HTMLAttributes, PropsWithChildren } from "react";

const SCROLL_BUFFER = 64;

const Navbar = ({ children }: PropsWithChildren) => {
  const { y: scrollY } = useScroll();

  const scrolled = scrollY > SCROLL_BUFFER;

  return (
    <nav
      className={cn(
        "fixed inset-0 z-40 flex h-nav w-full transition-[inset]",
        scrolled && "inset-y-2",
      )}
    >
      {/* BACKDROP */}
      <div
        className={cn(
          "w-full transition-[max-width,margin]",
          scrolled ? "mx-responsive" : "m-[0_auto] max-w-full",
        )}
      >
        <div
          className={cn(
            "inset-0 h-nav w-full rounded-[0px] bg-gradient-to-b from-black/80 outline outline-transparent backdrop-blur-none transition-[backdrop-filter,--tw-gradient-from,--tw-gradient-to,padding,border-radius,inset,outline-color] duration-200",
            scrolled &&
              "rounded-[2rem] from-zinc-950/75 to-zinc-950/75 outline-zinc-800 backdrop-blur-2xl",
          )}
        >
          <div
            className={cn(
              "inset-0 left-0 flex h-nav justify-between px-8 transition-[padding] duration-200",
              scrolled && "px-12",
            )}
          >
            {children}
          </div>
        </div>
      </div>
    </nav>
  );
};

const Header = ({ children }: PropsWithChildren) => {
  return <div className="flex items-center gap-6">{children}</div>;
};

const Content = ({ children }: PropsWithChildren) => {
  return <div className="flex items-center gap-8">{children}</div>;
};

type RouteEntryProps = {
  pathname: string;
  name: string;
};

const RouteList = ({
  children,
  className,
}: HTMLAttributes<HTMLUListElement>) => {
  return (
    <ul
      className={cn(
        "hidden items-center gap-8 text-lg text-zinc-50 md:flex",
        className,
      )}
    >
      {children}
    </ul>
  );
};

const RouteEntry = ({ pathname, name }: RouteEntryProps) => {
  const { pathname: pagePathname } = useLocation();

  return (
    <li>
      <Link
        to={pathname}
        title={name}
        className={cn(
          "text-text/75 transition-colors before:invisible before:block before:h-0 before:overflow-hidden before:font-semibold before:content-[attr(title)] hover:text-text [.is-current]:font-semibold [.is-current]:text-text",
          {
            "is-current": pathname === pagePathname,
          },
        )}
      >
        {name}
      </Link>
    </li>
  );
};

Navbar.Content = Content;
Navbar.Header = Header;
Navbar.RouteEntry = RouteEntry;
Navbar.RouteList = RouteList;

export default Navbar;
