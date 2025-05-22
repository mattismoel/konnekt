type Props = {
	title: string;
	description: string;
}

const PageMeta = ({ title, description }: Props) => (
	<>
		<title>{title}</title>
		<meta name="description" content={description} />
	</>
)

export default PageMeta
