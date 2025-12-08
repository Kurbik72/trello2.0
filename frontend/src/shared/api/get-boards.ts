import { httpService } from '../service/http-service'

interface GetBoardsResponse {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  backgroundSrc: string
}

const mockData: GetBoardsResponse[] = [
  {
    id: '1',
    title: 'Board 1',
    isFavorite: false,
    linkToBoard: 'https://board1.com',
    backgroundSrc: 'https://picsum.photos/id/100/800/400',
  },
  {
    id: '2',
    title: 'Board 2',
    isFavorite: false,
    linkToBoard: 'https://board2.com',
    backgroundSrc: 'https://picsum.photos/id/200/800/400',
  },
  {
    id: '3',
    title: 'Board 3',
    isFavorite: false,
    linkToBoard: 'https://board3.com',
    backgroundSrc: 'https://picsum.photos/id/400/800/400',
  },
]
export const getBoards = (user_id: string) =>
  httpService<GetBoardsResponse[]>(`/api/get-boards?user_id=${user_id}`, {
    method: 'GET',
    mock: {
      data: mockData,
      enabled: false,
    },
  })
