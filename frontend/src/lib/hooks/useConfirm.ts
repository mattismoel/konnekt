export const useConfirm = (msg: string, callback: () => void) => {
  if (!confirm(msg)) return

  callback()
}
