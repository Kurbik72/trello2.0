import { httpService } from '@/shared/service/http-service'

export interface DefaultBackground {
  id: string
  src: string
}

const mockData: DefaultBackground[] = [
  {
    id: '6949e844-9b4b-4dec-9d42-2d4216415f53',
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

export const getDefaultBackground = () =>
  httpService<DefaultBackground[]>('/api/default-backgrounds', {
    method: 'GET',
    mock: {
      data: mockData,
      enabled: true,
    },
  })
