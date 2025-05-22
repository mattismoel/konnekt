import StarterKit from "@tiptap/starter-kit"
import { EditorProvider, useCurrentEditor } from "@tiptap/react"
import Toolbar from "./toolbar";
import type { Level } from "@tiptap/extension-heading";
import { FaListOl, FaBold, FaItalic, FaListUl, FaUnderline } from "react-icons/fa";
import Underline from "@tiptap/extension-underline"

const HEADER_LEVELS: Level[] = [1, 2, 3]

const extensions = [StarterKit, Underline]

type Props = {
	content?: string | undefined;
	onChange?: (html: string) => void;
}

const Tiptap = ({ content, onChange }: Props) => {
	return (
		<div className="w-full">
			<EditorProvider
				content={content}
				extensions={extensions}
				slotBefore={<div><CustomToolbar /></div>}
				onTransaction={({ editor }) => onChange?.(editor.getHTML())}
				editorContainerProps={{ className: "max-w-none min-h-72 flex flex-col border border-zinc-900 prose prose-invert p-5 focus:outline-none rounded-b-md" }}
				editorProps={{
					attributes: { class: "focus:outline-none flex-1 w-full" }
				}}
			>
			</EditorProvider>
		</div>
	)
}

const CustomToolbar = () => {
	return (
		<Toolbar>
			<HeadingSection />
			<ListSection />
			<FormattingSection />
		</Toolbar>
	)
}

const HeadingSection = () => {
	const { editor } = useCurrentEditor()

	return (
		<Toolbar.ActionGroup>
			{HEADER_LEVELS.map(level => (
				<Toolbar.ActionGroup.Button
					key={level}
					title={`Heading ${level}`}
					active={editor?.isActive('heading', { level })}
					onClick={() => editor?.chain().focus().toggleHeading({ level }).run()}
				>
					H{level}
				</Toolbar.ActionGroup.Button>
			))}

		</Toolbar.ActionGroup>
	)
}

const ListSection = () => {
	const { editor } = useCurrentEditor()
	return (
		<Toolbar.ActionGroup>
			<Toolbar.ActionGroup.Button
				disabled={!editor?.isEditable}
				title="Bullet List"
				active={editor?.isActive('bulletList')}
				onClick={() => editor?.chain().focus().toggleBulletList().run()}
			>
				<FaListUl />
			</Toolbar.ActionGroup.Button>
			<Toolbar.ActionGroup.Button
				disabled={!editor?.isEditable}
				title="Numbered List"
				active={editor?.isActive('orderedList')}
				onClick={() => editor?.chain().focus().toggleOrderedList().run()}
			>
				<FaListOl />
			</Toolbar.ActionGroup.Button>
		</Toolbar.ActionGroup>
	)
}


const FormattingSection = () => {
	const { editor } = useCurrentEditor()

	return (
		<Toolbar.ActionGroup>
			<Toolbar.ActionGroup.Button
				disabled={!editor?.isEditable}
				title="Toggle Bold"
				active={editor?.isActive('bold')}
				onClick={() => editor?.chain().focus().toggleBold().run()}
			>
				<FaBold />
			</Toolbar.ActionGroup.Button>
			<Toolbar.ActionGroup.Button
				disabled={!editor?.isEditable}
				title="Toggle Italic"
				active={editor?.isActive('italic')}
				onClick={() => editor?.chain().focus().toggleItalic().run()}
			>
				<FaItalic />
			</Toolbar.ActionGroup.Button>
			<Toolbar.ActionGroup.Button
				disabled={!editor?.isEditable}
				title="Toggle Underline"
				active={editor?.isActive('underline')}
				onClick={() => editor?.chain().focus().toggleUnderline().run()}
			>
				<FaUnderline />
			</Toolbar.ActionGroup.Button>
		</Toolbar.ActionGroup>
	)
}

export default Tiptap
