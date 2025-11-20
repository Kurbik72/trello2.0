import { useFetch } from '@vueuse/core'
import { ref } from 'vue'

export interface DefaultBackground {
  id: string
  src: string
}

const mockData: DefaultBackground[] = [
  {
    id: '1',
    src: 'https://picsum.photos/id/1/800/400',
  },
  {
    id: '2',
    src: 'https://picsum.photos/id/2/800/400',
  },
  {
    id: '3',
    src: 'https://picsum.photos/id/3/800/400',
  },
]

const IS_USE_MOCK = true

async function getMockResponse(): Promise<DefaultBackground[]> {
  await new Promise((resolve) => {
    setTimeout(resolve, 500)
  })

  return mockData
}

export const getDefaultBackground = () => {
  if (IS_USE_MOCK) {
    const data = ref<DefaultBackground[] | null>(null)
    const isFetching = ref(true)
    const error = ref<Error | null>(null)

    getMockResponse()
      .then((result) => {
        data.value = result
        isFetching.value = false
      })
      .catch((err) => {
        error.value = err
        isFetching.value = false
      })

    return {
      data,
      isFetching,
      error,
      execute: async () => {
        isFetching.value = true
        error.value = null
        try {
          data.value = await getMockResponse()
        } catch (err) {
          error.value = err as Error
        } finally {
          isFetching.value = false
        }
      },
    }
  }

  return useFetch<DefaultBackground[]>('/api/default-backgrounds', {
    method: 'GET',
  })
}
