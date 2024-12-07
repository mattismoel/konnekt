import type { Stream } from "stream";

const extensionRegex = /(?:\.([^.]+))?$/;

/**
 * @description Finds the extension of an input filename.
 * @param {string} filename - The filename to find the extension of.
 * @returns The extension of the input filename.
 */
export const fileExtension = (filename: string): string => {
  const result = extensionRegex.exec(filename)

  if (!result) {
    throw Error("No extension found")
  }

  return result[1]
}

/**
 * @description Reads a stream and returns it as a buffer.
 */
export const streamToBuffer = async (stream: Stream): Promise<Buffer> => {
  return new Promise<Buffer>((resolve, reject) => {
    const buf = Array<any>();
    stream.on("data", (chunk) => buf.push(chunk))
    stream.on("end", () => resolve(Buffer.concat(buf)))
    stream.on("error", (e) => reject(e))
  })
}
