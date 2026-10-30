import type { ApiResponse } from '../types/api'

const API_URL = '/api'

type Method = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE'

interface RequestOptions {
  method?: Method
  body?: unknown
}

export class ApiError extends Error {
  status: number

  constructor(status: number, message: string) {
    super(message)
    this.status = status
  }
}

async function send<T>(
  path: string,
  options: RequestOptions
): Promise<ApiResponse<T>> {
  const { method = 'GET', body } = options
  const isForm = body instanceof FormData
  let res: Response
  try {
    res = await fetch(`${API_URL}${path}`, {
      method,
      credentials: 'include',
      headers:
        body === undefined || isForm
          ? undefined
          : { 'Content-Type': 'application/json' },
      body:
        body === undefined ? undefined : isForm ? body : JSON.stringify(body),
    })
  } catch {
    throw new ApiError(0, 'Could not reach the server')
  }

  let json: unknown = null
  try {
    json = await res.json()
  } catch {}

  if (res.ok) return json as ApiResponse<T>

  const fail = json as {
    message?: string
    meta?: { errorCode?: string }
  } | null
  throw new ApiError(res.status, fail?.message ?? 'Something went wrong')
}

export async function apiRequest<T>(
  path: string,
  options: RequestOptions = {}
): Promise<ApiResponse<T>> {
  try {
    return await send<T>(path, options)
  } catch (err) {
    throw err
  }
}
