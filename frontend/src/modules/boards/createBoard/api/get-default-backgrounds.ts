import { httpService } from '@/shared/service/http-service'

export interface DefaultBackground {
  id: string
  src: string
}

const mockData: DefaultBackground[] = [
  {
    id: '6949e844-9b4b-4dec-9d42-2d4216415f53',
    src: 'https://picsum.photos/id/100/1920/1080',
  },
  {
    id: '2',
    src: 'https://picsum.photos/id/200/1920/1080',
  },
  {
    id: '4',
    src: 'https://picsum.photos/id/300/1920/1080',
  },
  {
    id: '5',
    src: 'https://picsum.photos/id/400/1920/1080',
  },
  {
    id: '6',
    src: 'https://picsum.photos/id/500/1920/1080',
  },
  {
    id: '7',
    src: 'https://picsum.photos/id/600/1920/1080',
  },
  {
    id: '8',
    src: 'https://picsum.photos/id/900/1920/1080',
  },
]

export const getDefaultBackground = () =>
  httpService<DefaultBackground[]>('/api/default-backgrounds', {
    method: 'GET',
    mock: {
      data: mockData,
      enabled: false,
    },
  })
