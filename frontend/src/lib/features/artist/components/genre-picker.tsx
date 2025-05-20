import type { Genre } from "../features/artist/genre";
import Button from "./ui/button/button";
import Modal from "./ui/modal"
import Picker from "./ui/picker";

type Props = {
	show: boolean;
	onClose: () => void;

	genres: Genre[]
}

const GenrePicker = ({ show, onClose }: Props) => {
	return (
		<Picker show={show} onClose={onClose}>

		</Picker>
	)
}

export default GenrePicker
