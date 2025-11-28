type Src = {
  src: string;
  alt: string;
};

type Props = {
  srcs: Map<string, Src>;
};

const SponsorDisplay = ({ srcs }: Props) => {
  return (
    <div className="flex flex-wrap justify-center gap-12 bg-radial from-zinc-900 to-background to-[75%] p-8 sm:p-16">
      {Array.from(srcs).map(([href, { src, alt }]) => (
        <a
          key={href}
          aria-label={alt}
          href={href}
          className="group flex w-20 items-center justify-center sm:w-24"
        >
          <img
            src={src}
            alt={alt}
            loading="lazy"
            className="aspect-square transition-[filter,scale] group-hover:brightness-100 md:brightness-60"
          />
        </a>
      ))}
    </div>
  );
};

export default SponsorDisplay;
