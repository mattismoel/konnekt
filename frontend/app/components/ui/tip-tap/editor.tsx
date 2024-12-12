import { EditorContent, useEditor } from "@tiptap/react";
import StarterKit from "@tiptap/starter-kit";
import Underline from "@tiptap/extension-underline";
import Link from "@tiptap/extension-link";
import YouTube from "@tiptap/extension-youtube";
import { BulletOptions, HeadingOptions, HistoryOptions, MediaOptions, StyleOptions } from "./options";

type Props = {
  content?: string;
  onChange?: (value: string) => void;
}

export const TipTapEditor = ({ content = "", onChange }: Props) => {
  const editor = useEditor({
    extensions: [StarterKit.configure({
      heading: { levels: [1, 2, 3] },
    }),
      Underline,
    Link.configure({
      openOnClick: true,
      autolink: true,
    }),
    YouTube.configure({
      HTMLAttributes: {
        class: "w-full aspect-video",
      },
    }),
    ],
    editorProps: {
      attributes: {
        class: "prose prose-invert focus:outline-none h-full"
      }
    },
    content: content,
    onUpdate: ({ editor }) => {
      onChange?.(editor.getHTML())
    },
    immediatelyRender: false,
  })

  return (
    <>
      <div>
        {editor && (
          <div
            className="w-full overflow-x-scroll px-4 bg-neutral-900 
            text-neutral-500 h-12 flex rounded-t-sm border border-neutral-800 
            divide-x divide-neutral-800 *:px-6 first:*:pl-0 last:*:pr-0"
          >
            <HistoryOptions editor={editor} />
            <HeadingOptions editor={editor} />
            <BulletOptions editor={editor} />
            <StyleOptions editor={editor} />
            <MediaOptions editor={editor} />
          </div >
        )}
        <div className="h-96 w-full p-4 bg-neutral-950 border-x border-b 
          border-neutral-900 rounded-b-sm overflow-scroll">
          <EditorContent className="h-full" editor={editor} />
        </div>
      </div >
    </>
  )
}
