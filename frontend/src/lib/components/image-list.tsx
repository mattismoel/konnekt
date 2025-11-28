import { FaXmark } from "react-icons/fa6";
import type { Image } from "../features/content/landing";
import type { ID } from "../api";
import { useAuth } from "../context/auth";

type Props = {
  images: Image[];

  onDelete: (id: ID) => void;
};

const ImageList = ({ images, onDelete }: Props) => {
  const { hasPermissions } = useAuth();
  const isEditable = hasPermissions(["edit:content", "delete:content"]);

  return (
    <div className="flex flex-wrap gap-8">
      {images.map((img) => (
        <div className="relative h-48">
          <img
            src={img.url}
            className="h-full rounded-sm border border-zinc-800"
          />
          {isEditable && (
            <button
              type="button"
              className="absolute top-0 right-0 translate-x-1/2 -translate-y-1/2 rounded-full border border-zinc-700 bg-zinc-800 p-1 text-text hover:bg-zinc-700"
              onClick={() => onDelete(img.id)}
            >
              <FaXmark />
            </button>
          )}
        </div>
      ))}
    </div>
  );
};

export default ImageList;
