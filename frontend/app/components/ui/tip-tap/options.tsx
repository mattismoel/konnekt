import { useEffect, useRef, useState } from "react"
import { Editor } from "@tiptap/react"
import { BiListUl, BiListOl, BiBold, BiItalic, BiUnderline, BiLink, BiLogoYoutube, BiUndo, BiRedo } from "react-icons/bi"
import { Input } from "../input"
import { Button } from "../button"

type Props = {
  editor: Editor
}

export const BulletOptions = ({ editor }: Props) => {
  return (
    <div className="flex gap-4 items-center">
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleBulletList().run()}
        className={`${editor.isActive("bulletList") && "text-neutral-100"}`}
      >
        <BiListUl />
      </button>
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleOrderedList().run()}
        className={`${editor.isActive("orderedList") && "text-neutral-100"}`}
      >
        <BiListOl />
      </button>
    </div>
  )
}

export const HeadingOptions = ({ editor }: Props) => {
  return (
    <div className="flex gap-4 items-center">
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleHeading({ level: 1 }).run()}
        className={`${editor.isActive("heading", { level: 1 }) && "text-neutral-100 font-bold"}`}
      >H1</button
      >
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleHeading({ level: 2 }).run()}
        className={`${editor.isActive("heading", { level: 2 }) && "text-neutral-100 font-bold"}`}
      >H2</button
      >
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleHeading({ level: 3 }).run()}
        className={`${editor.isActive("heading", { level: 3 }) && "text-neutral-100 font-bold"}`}
      >H3</button
      >
    </div>

  )
}

export const HistoryOptions = ({ editor }: Props) => {
  return (
    <div className="flex items-center gap-4">
      <button
        type="button"
        onClick={() => editor.chain().focus().undo().run()}
        disabled={!editor.can().undo()}
        className={`${editor.can().undo() && "text-neutral-100"}`}
      >
        <BiUndo />
      </button>
      <button
        type="button"
        onClick={() => editor.chain().focus().redo().run()}
        disabled={!editor.can().redo()}
        className={`${editor.can().redo() && "text-neutral-100"}`}
      >
        <BiRedo />
      </button>
    </div>
  )
}


export const MediaOptions = ({ editor }: Props) => {
  const addLinkModal = useRef<HTMLDialogElement>(null)
  const youtubeModal = useRef<HTMLDialogElement>(null)

  let [showLinkModal, setShowLinkModal] = useState(false);
  let [showYoutubeModal, setShowYoutubeModal] = useState(false);

  let [linkUrl, setLinkUrl] = useState("");
  let [youTubeUrl, setYoutubeUrl] = useState("");

  const setLink = () => {
    setShowLinkModal(false);

    if (linkUrl === "") {
      editor.chain().focus().extendMarkRange("link").unsetLink().run();
      return;
    }

    editor
      .chain()
      .focus()
      .extendMarkRange("link")
      .setLink({ href: linkUrl })
      .run();
  };

  const setYouTubeVideo = () => {
    showYoutubeModal = false;
    if (!editor) return;

    if (!youTubeUrl) return;

    editor.commands.setYoutubeVideo({
      src: youTubeUrl,
    });
  };



  useEffect(() => {
    if (showLinkModal) addLinkModal.current?.showModal();
    else addLinkModal.current?.close();
  });

  useEffect(() => {
    if (showYoutubeModal) youtubeModal.current?.showModal();
    else youtubeModal.current?.close();
  }, [showYoutubeModal]);


  return (

    <>
      <div className="flex items-center gap-4 text-neutral-100">
        <button type="button" onClick={() => setShowLinkModal(true)}
        >< BiLink /></button
        >
        <button type="button" onClick={() => setShowYoutubeModal(true)}
        ><BiLogoYoutube /></button
        >
      </div>

      <div>
        <dialog
          ref={addLinkModal}
          onClose={() => (setShowLinkModal(false))}
          className="p-6 bg-neutral-900 text-neutral-100 border border-neutral-700 rounded-sm"
        >
          <h1 className="text-center mb-4 font-semibold">Tilføj link</h1>
          <div className="space-y-2">
            <Input name="" placeholder="Link" onChange={(e) => setLinkUrl(e.currentTarget.value)} />
            <Button type="button" onClick={setLink}>Tilføj</Button>
          </div>
        </dialog>

        <div>
          <dialog
            ref={youtubeModal}
            onClose={() => setShowYoutubeModal(false)}
            className="p-6 bg-neutral-900 text-white border border-neutral-700 rounded-sm"
          >
            <h1 className="text-center mb-4 font-semibold">Tilføj Youtube-video</h1>
            <div className="space-y-2">
              <Input name="" placeholder="Link" onChange={(e) => setYoutubeUrl(e.currentTarget.value)} />
              <Button type="button" onClick={setYouTubeVideo}>Tilføj</Button>
            </div>
          </dialog>
        </div>
      </div>
    </>
  )
}

export const StyleOptions = ({ editor }: Props) => {
  return (
    <div className="flex gap-4 items-center">
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleBold().run()}
        className={`${editor.isActive("bold") && "text-neutral-100"}`}
      ><BiBold /></button
      >
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleItalic().run()}
        className={`${editor.isActive("italic") && "text-neutral-100"}`}
      ><BiItalic /></button
      >
      <button
        type="button"
        onClick={() => editor.chain().focus().toggleUnderline().run()}
        className={`${editor.isActive("underline") && "text-neutral-100"}`}
      ><BiUnderline /></button
      >
    </div>
  )
}
