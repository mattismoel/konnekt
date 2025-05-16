import type { IconType } from 'react-icons/lib';
import { cn } from '../clsx';
import type { Severity, Toast as ToastType } from '../context/toast';
import { FaXmark } from "react-icons/fa6"
import { FaExclamationCircle, FaExclamationTriangle, FaInfoCircle } from 'react-icons/fa';

let iconMap = new Map<Severity, IconType>([
	['warning', FaExclamationTriangle],
	['info', FaInfoCircle],
	['error', FaExclamationCircle]
]);


type Props = {
	toast: ToastType;
	onDelete: () => void;
};

const Toast = ({ toast, onDelete }: Props) => {
	const Icon = iconMap.get(toast.severity)!

	return (
		<div
			className={cn('animate-fade-in pointer-events-none relative flex min-w-sm flex-col rounded-sm border p-2', {
				'border-blue-900 bg-blue-950': toast.severity === 'info',
				'border-red-900 bg-red-950': toast.severity === 'error',
				'border-yellow-900 bg-yellow-950': toast.severity === 'warning'
			})}
		>
			<div
				className={cn("flex items-center gap-2", {
					"text-blue-200": toast.severity === "info",
					"text-yellow-200": toast.severity == "warning",
					"text-red-200": toast.severity === "error",
				})}
			>
				<Icon />
				<span className={cn({ "font-medium": toast.message && toast.message !== "" })}>
					{toast.title}
				</span>
				<button
					type="button"
					onClick={onDelete}
					className={cn("pointer-events-auto absolute top-2 right-2 rounded-full p-1", {
						'border-blue-900 bg-blue-950': toast.severity === 'info',
						'border-red-900 bg-red-950': toast.severity === 'error',
						'border-yellow-900 bg-yellow-950': toast.severity === 'warning'
					})}
				>
					<FaXmark />
				</button>
			</div>
			<span
				className={cn("whitespace-pre-wrap", {
					"text-blue-300": toast.severity === 'info',
					"text-yellow-300": toast.severity === 'warning',
					"text-red-300": toast.severity === 'error'
				})}
			>
				{toast.message}
			</span>
		</div>
	)
}

export default Toast
