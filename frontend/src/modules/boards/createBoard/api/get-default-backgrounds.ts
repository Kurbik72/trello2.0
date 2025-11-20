import { useFetch } from '@vueuse/core'
import { ref } from 'vue'
import defaultBackground1 from '@/assets/images/2.png'
import defaultBackground2 from '@/assets/images/3.png'
import defaultBackground3 from '@/assets/images/4.png'

export interface DefaultBackground {
  id: string
  src: string
}

const mockData: DefaultBackground[] = [
  {
    id: '1',
    src: defaultBackground1,
  },
  {
    id: '2',
    src: defaultBackground2,
  },
  {
    id: '3',
    src: defaultBackground3,
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
