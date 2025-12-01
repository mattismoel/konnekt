import { cn } from "@/lib/clsx";
import type { HTMLAttributes } from "react";

const Card = ({
  children,
  className,
  ...rest
}: HTMLAttributes<HTMLDivElement>) => {
  return (
    <div
      {...rest}
      className={cn("rounded-xl border border-zinc-800 bg-zinc-950", className)}
    >
      {children}
    </div>
  );
};

const Header = ({
  children,
  className,
  ...rest
}: HTMLAttributes<HTMLDivElement>) => {
  return (
    <div {...rest} className={cn("flex flex-col gap-2 p-8 pt-12", className)}>
      {children}
    </div>
  );
};

const Content = ({
  children,
  className,
  ...rest
}: HTMLAttributes<HTMLDivElement>) => {
  return (
    <div {...rest} className={cn("flex flex-col p-6", className)}>
      {children}
    </div>
  );
};

const Title = ({
  children,
  className,
  ...rest
}: HTMLAttributes<HTMLHeadingElement>) => {
  return (
    <h1
      {...rest}
      className={cn(
        "mb-4 font-heading text-3xl font-bold text-text-light",
        className,
      )}
    >
      {children}
    </h1>
  );
};

const Description = ({
  children,
  ...rest
}: HTMLAttributes<HTMLParagraphElement>) => {
  return (
    <p {...rest} className="text-text/50">
      {children}
    </p>
  );
};

const Footer = ({ children, ...rest }: HTMLAttributes<HTMLDivElement>) => {
  return (
    <div {...rest} className="p-6 pb-12">
      {children}
    </div>
  );
};

Card.Header = Header;
Card.Content = Content;

Card.Title = Title;
Card.Description = Description;
Card.Footer = Footer;

export default Card;
