import { useHttpService } from '@/shared/service/http-service'

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

export const getDefaultBackground = () => {
  return useHttpService<DefaultBackground[]>('/api/default-backgrounds', {
    mock: {
      enabled: true,
      data: mockData,
      delay: 1000,
    },
  })
}
