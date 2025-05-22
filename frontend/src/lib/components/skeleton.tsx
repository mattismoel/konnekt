type SkeletonTextProps = {
	wordCount?: number
}

export const SkeletonText = ({ wordCount = 1 }: SkeletonTextProps) => (
	<div className="animate-pulse flex gap-2">
		{[...Array(wordCount)].map((_, index) => (
			<div
				key={index}
				className="shrink-0 rounded-full h-4 bg-text/75"
				style={{ width: `${32 + Math.random() * 64}px` }}
			/>
		))}
	</div>
)

export const SkeletonIcon = () => (
	<div className="animate-pulse h-6 w-6 rounded-md bg-text/75" />
)
