import { z, ZodSchema } from "zod"
import { useToast } from "./context/toast";
import { useNavigate } from "@tanstack/react-router";

export const idSchema = z
  .number()
  .int()
  .positive()

export type ID = z.infer<typeof idSchema>

export const apiErrorSchema = z.object({
  message: z.string(),
})

export class APIError extends Error {
  public readonly cause: string;
  public readonly status: number;

  constructor(status: number, message: string, cause: string) {
    super(message)
    this.name = "APIError"
    this.status = status
    this.cause = cause
    Object.setPrototypeOf(this, APIError.prototype)
  };
}

type Method = "GET" | "POST" | "PUT" | "PATCH" | "DELETE"

type RequestBody<T> = {
  bodySchema: Zod.Schema<T>
  body: T,
}

export async function requestAndParse<TBody, TResponse>(
  url: string | URL,
  resSchema: ZodSchema<TResponse>,
  errorMsg?: string,
  bodyOpts?: RequestBody<TBody>,
  method?: Method
): Promise<TResponse>

export async function requestAndParse<TBody>(
  url: string | URL,
  resSchema: undefined,
  errorMsg?: string,
  bodyOpts?: RequestBody<TBody>,
  method?: Method
): Promise<void>;

export async function requestAndParse<TBody, TResponse>(
  url: string | URL,
  resSchema?: Zod.Schema<TResponse>,
  errorMsg?: string,
  bodyOpts?: RequestBody<TBody>,
  method?: Method,
): Promise<TResponse | void> {
  const res = await fetch(url, {
    headers: {
      "Content-Type": "application/json"
    },
    credentials: "include",
    method,
    body: bodyOpts
      ? JSON.stringify(bodyOpts.bodySchema.parse(bodyOpts.body))
      : undefined
  })

  if (!isResponseSuccessful(res)) {
    const err = apiErrorSchema.parse(await res.json())

    console.dir(errorMsg + " " + err.message, { depth: Infinity })

    throw new APIError(
      res.status,
      errorMsg || "Something went wrong...",
      err.message,
    )
  }

  if (resSchema === undefined) return

  const data = resSchema.parse(await res.json())

  return data
}

const isResponseSuccessful = (res: Response): boolean => {
  return (res.status >= 200 && res.status < 300)
}

type SubmitHandlerOptions<T> = {
  action: (form: T) => Promise<any>
  successMessage: string;
  errorMessage?: string;
  navigateTo?: string;
}

export const createSubmitHandler = <T>(
  options: SubmitHandlerOptions<T>
): (form: T) => Promise<void> => {
  const {
    action,
    successMessage,
    errorMessage = "Noget gik galt...",
    navigateTo = "/admin/events",
  } = options

  const { addToast } = useToast()
  const navigate = useNavigate()

  return async (form: T) => {
    try {
      await action(form)
      addToast(successMessage)
      navigate({ to: navigateTo })
    } catch (e) {
      if (e instanceof APIError) {
        addToast(errorMessage, e.message, "error")
        return
      }

      addToast(errorMessage, "Noget gik galt...", "error")
      throw e
    }
  }
}
