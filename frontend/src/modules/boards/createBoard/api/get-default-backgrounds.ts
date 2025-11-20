import { useFetch } from '@vueuse/core'
import { ref } from 'vue'

export interface DefaultBackground {
  id: string
  src: string
}

const mockData: DefaultBackground[] = [
  {
    id: '1',
    src: 'https://picsum.photos/id/100/800/400',
  },
  {
    id: '2',
    src: 'https://picsum.photos/id/200/800/400',
  },
  {
    id: '4',
    src: 'https://picsum.photos/id/300/800/400',
  },
  {
    id: '5',
    src: 'https://picsum.photos/id/400/800/400',
  },
  {
    id: '6',
    src: 'https://picsum.photos/id/500/800/400',
  },
  {
    id: '7',
    src: 'https://picsum.photos/id/600/800/400',
  },
  {
    id: '8',
    src: 'https://picsum.photos/id/700/800/400',
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
