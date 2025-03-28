<script lang="ts">
	import { Editor } from '@tiptap/core';

	import Document from '@tiptap/extension-document';
	import Paragraph from '@tiptap/extension-paragraph';
	import Text from '@tiptap/extension-text';
	import { Heading, type Level } from '@tiptap/extension-heading';
	import Placeholder from '@tiptap/extension-placeholder';
	import Bold from '@tiptap/extension-bold';
	import Italic from '@tiptap/extension-italic';
	import Underline from '@tiptap/extension-underline';

	import BulletList from '@tiptap/extension-bullet-list';
	import OrderedList from '@tiptap/extension-ordered-list';
	import ListItem from '@tiptap/extension-list-item';

	import ActionGroup from './ActionGroup.svelte';
	import ActionButton from './ActionButton.svelte';
	import Toolbar from './Toolbar.svelte';

	import ListBulletedIcon from '~icons/mdi/format-list-bulleted';
	import ListOrderedIcon from '~icons/mdi/format-list-numbered';

	import BoldIcon from '~icons/mdi/format-bold';
	import ItalicIcon from '~icons/mdi/format-italic';
	import UnderlineIcon from '~icons/mdi/format-underline';

	import Content from './Content.svelte';
	import { onMount } from 'svelte';

	const HEADER_LEVELS: Level[] = [1, 2, 3];

	type Props = {
		value: string;
	};

	let { value = $bindable() }: Props = $props();

	let element: HTMLDivElement;
	let editor = $state<Editor | null>(null);

	onMount(() => {
		editor = new Editor({
			element,
			extensions: [
				Document,
				Paragraph,
				Text,
				Bold,
				Italic,
				Underline,
				Placeholder.configure({
					placeholder: 'Eventbeskrivelse...',
					emptyNodeClass:
						'cursor-text before:content-[attr(data-placeholder)] before:absolute before:text-text/50 before:pointer-events-none'
				}),
				Heading.configure({ levels: HEADER_LEVELS }),
				BulletList,
				ListItem,
				OrderedList
			],
			content: value,
			editorProps: {
				attributes: {
					class: 'prose prose-invert m-5 focus:outline-none'
				}
			},
			onTransaction: () => {
				editor = editor;
			},
			onUpdate: ({ editor }) => {
				value = editor.getHTML();
			}
		});

		return () => {
			if (!editor) return;
			editor.destroy();
		};
	});
</script>

<div class="flex flex-col">
	{#if editor}
		<Toolbar>
			<ActionGroup>
				{#each HEADER_LEVELS as level}
					<ActionButton
						title="Heading {level}"
						active={editor.isActive('heading', { level })}
						onclick={() => editor?.chain().focus().toggleHeading({ level }).run()}
					>
						H{level}
					</ActionButton>
				{/each}
			</ActionGroup>
			<ActionGroup>
				<ActionButton
					title="Bullet List"
					active={editor.isActive('bulletList')}
					onclick={() => editor?.chain().focus().toggleBulletList().run()}
				>
					<ListBulletedIcon />
				</ActionButton>
				<ActionButton
					title="Numbered List"
					active={editor.isActive('orderedList')}
					onclick={() => editor?.chain().focus().toggleOrderedList().run()}
				>
					<ListOrderedIcon />
				</ActionButton>
			</ActionGroup>
			<div class="flex-1"></div>
			<ActionGroup>
				<ActionButton
					title="Toggle Bold"
					active={editor.isActive('bold')}
					onclick={() => editor?.chain().focus().toggleBold().run()}
				>
					<BoldIcon />
				</ActionButton>
				<ActionButton
					title="Toggle Italic"
					active={editor.isActive('italic')}
					onclick={() => editor?.chain().focus().toggleItalic().run()}
				>
					<ItalicIcon />
				</ActionButton>
				<ActionButton
					title="Toggle Underline"
					active={editor.isActive('underline')}
					onclick={() => editor?.chain().focus().toggleUnderline().run()}
				>
					<UnderlineIcon />
				</ActionButton>
			</ActionGroup>
		</Toolbar>
	{/if}
	<Content>
		<div bind:this={element}></div>
	</Content>
</div>
