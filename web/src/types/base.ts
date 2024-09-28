export interface BaseResponse<T = null> {
  status: 'success' | 'error'
  message: string
  data: T
}
