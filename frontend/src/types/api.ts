export interface ApiResponse<T> {
  message: string
  meta: { code: number; status: string }
  data: T
}
