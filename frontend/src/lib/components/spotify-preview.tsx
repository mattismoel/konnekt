type Props = {
		trackId: string;
}

const SpotifyPreview = ({trackId}: Props) => {
	let src = 
		`https://open.spotify.com/embed/track/${trackId}?utm_source=generator&theme=0`

	return (
		<div className="max-w-2xl">
			<iframe
				src={src}
				title="Audio preview"
				style:border-radius="12px"
				width="100%"
				height="152"
				frameBorder="0"
				allowFullScreen
				allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture"
				loading="lazy"
			></iframe>
		</div>
	)
}

export default SpotifyPreview
