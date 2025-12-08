import axios, { type AxiosRequestConfig } from 'axios'
import { ref, type Ref } from 'vue'

interface MockError {
  errorMessage: string
  statusCode: number
}

interface MockOptions<ResponseType> {
  data: ResponseType
  isError?: boolean
  error?: MockError
  delay?: number
  enabled?: boolean
}

interface RequestOptions<ResponseType, BodyType = unknown> extends AxiosRequestConfig {
  mock?: MockOptions<ResponseType>
  body?: BodyType
}

export interface HttpResponse<ResponseType> {
  data: Ref<ResponseType | null>
  error: Ref<MockError | null>
  isLoading: Ref<boolean>
  execute: () => Promise<void>
}

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_HTTP_API_PATH,
  headers: {
    'Content-Type': 'application/json',
  },
})

export const httpService = <ResponseType, BodyType = unknown>(
  url: string,
  options: RequestOptions<ResponseType, BodyType>,
): HttpResponse<ResponseType> => {
  const data = ref<ResponseType | null>(null) as Ref<ResponseType | null>
  const error = ref<MockError | null>(null)
  const isLoading = ref(false)

  const execute = async () => {
    isLoading.value = true
    error.value = null

    try {
      if (options.mock?.enabled) {
        const mockDelay = options.mock.delay ?? 1000

        await new Promise((resolve) => setTimeout(resolve, mockDelay))

        if (options.mock.isError) {
          error.value = {
            errorMessage: options.mock.error?.errorMessage || 'Internal Server Error',
            statusCode: options.mock.error?.statusCode || 500,
          }
          return
        }

        data.value = options.mock.data
        return
      }

      // Преобразуем body в data для axios
      // eslint-disable-next-line @typescript-eslint/no-unused-vars
      const { body, mock, ...axiosConfig } = options
      const response = await apiClient.request<ResponseType>({
        url,
        ...axiosConfig,
        data: body, // axios использует data вместо body
      })
      data.value = response.data
    } catch (err) {
      if (axios.isAxiosError(err)) {
        error.value = {
          errorMessage: err.message,
          statusCode: err.response?.status || 500,
        }
      } else {
        error.value = {
          errorMessage: 'Unknown error',
          statusCode: 500,
        }
      }
    } finally {
      isLoading.value = false
    }
  }

  return {
    data,
    error,
    isLoading,
    execute,
  }
}
