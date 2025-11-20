import { ref, computed } from 'vue'

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

interface RequestOptions<DataType> {
  mock?: MockOptions<DataType>
}

export const useHttpService = <DataType>(url: string, options: RequestOptions<DataType>) => {
  // Mock режим
  if (options.mock?.enabled) {
    const isFetching = ref(true)
    const error = ref<Error | null>(null)
    const data = ref<DataType | null>(null)

    const execute = async () => {
      isFetching.value = true
      error.value = null

      const mockDelay = options.mock!.delay || 1000
      await new Promise((resolve) => setTimeout(resolve, mockDelay))

      if (options.mock!.error) {
        error.value = new Error(options.mock!.error.errorMessage)
      } else {
        data.value = options.mock!.data
      }
      isFetching.value = false
    }

    // Автоматически выполняем запрос при создании
    execute()

    return {
      data: computed(() => data.value),
      error: computed(() => error.value),
      isFetching: computed(() => isFetching.value),
      execute,
    }
  }

  // Реальный режим
  const { mock, ...fetchOptions } = options
  return useFetch<DataType>(url, fetchOptions).json()
}
