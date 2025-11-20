import { useFetch, type UseFetchOptions } from '@vueuse/core'

interface MockError {
  errorMessage: string
  statusCode: number
}

interface MockOptions<DataType> {
  data: DataType
  isError?: boolean
  error?: MockError
  delay?: number
  enabled?: boolean
}

interface RequestOptions<DataType> extends UseFetchOptions {
  mock: MockOptions<DataType>
}

export const httpService = async <DataType>(url: string, options: RequestOptions<DataType>) => {
  if (options.mock.enabled) {
    const mockDelay = options.mock.delay || 1000

    await new Promise((resolve) => setTimeout(resolve, mockDelay))

    if (options.mock.error) {
      throw {
        errorMessage: options.mock.error?.errorMessage || 'Internal Server Error',
        statusCode: options.mock.error?.statusCode || 500,
      } as MockError
    }

    return options.mock.data
  }

  return useFetch(url, options)
}
