import Logo from "@/lib/assets/logo";
import { Link } from "@tanstack/react-router";
import { FaFacebook, FaInstagram, FaTiktok } from "react-icons/fa6";
import type { IconType } from "react-icons/lib";

type Src = {
  icon: IconType;
  title: string;
};

type ContactType = "phone" | "mail";
type NavPath = string;

const contactTypeMap = new Map<ContactType, string>([
  ["mail", "mailto"],
  ["phone", "tel"],
]);

const socialMap = new Map<string, Src>([
  [
    "https://www.instagram.com/konnekt_odense/",
    { icon: FaInstagram, title: "Instagram" },
  ],
  ["https://www.tiktok.com/@konnekt_", { icon: FaTiktok, title: "TikTok" }],
  [
    "https://www.facebook.com/profile.php?id=61574860865073",
    { icon: FaFacebook, title: "Facebook" },
  ],
]);

const navMap = new Map<NavPath, string>([
  ["/", "Hjem"],
  ["/events", "Events"],
  ["/artists", "Kunstnere"],
  ["/about", "Om os"],
]);

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="border-t border-t-zinc-900 bg-zinc-950 pt-8 pb-6">
      <div className="mx-responsive">
        <div className="mb-8 grid w-full grid-cols-1 gap-8 sm:grid-cols-2">
          <div className="flex-1">
            <span className="mb-2 inline-block align-top font-heading leading-none font-bold text-text-light">
              Find rundt
            </span>
            <NavList links={navMap} />
          </div>

          <div className="flex flex-1 flex-col items-start gap-4 sm:items-end">
            <Logo className="mb-2 hidden h-4 fill-text-light sm:block" />
            <span className="mb-2 font-heading font-bold sm:hidden">
              Kontakt os
            </span>

            <address className="flex flex-col items-start gap-2 not-italic sm:items-end">
              <ContactEntry type="mail" value="konnekt.samarbejde@gmail.com" />
              <ContactEntry type="mail" value="booking.konnekt@gmail.com" />
            </address>

            <SocialMediaList socialMap={socialMap} />
          </div>
        </div>

        <span className="line-clamp-1 text-center text-xs">
          &copy; {currentYear} Foreningen Konnekt
        </span>
      </div>
    </footer>
  );
};

const SocialMediaList = ({ socialMap }: { socialMap: Map<string, Src> }) => (
  <ul className="flex items-center gap-4 text-xl">
    {Array.from(socialMap).map(([href, { icon: Icon, title }]) => (
      <li key={href} className="hover:text-text-light">
        <a title={title} href={href} target="_blank">
          <Icon />
        </a>
      </li>
    ))}
  </ul>
);

const ContactEntry = ({
  value,
  type,
}: {
  type: ContactType;
  value: string;
}) => (
  <a
    href={`${contactTypeMap.get(type)}:${value}`}
    className="transition-colors hover:text-text-light hover:underline"
  >
    {value}
  </a>
);

const NavList = ({ links }: { links: Map<NavPath, string> }) => (
  <ul>
    {Array.from(links).map(([path, name]) => (
      <li className="transition-colors hover:text-text-light hover:underline">
        <Link to={path}>{name}</Link>
      </li>
    ))}
  </ul>
);

export default Footer;
