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
    backgroundSrc: 'https://board1.com/background.png',
  },
  {
    id: '2',
    title: 'Board 2',
    isFavorite: false,
    linkToBoard: 'https://board2.com',
    backgroundSrc: 'https://board2.com/background.png',
  },
  {
    id: '3',
    title: 'Board 3',
    isFavorite: false,
    linkToBoard: 'https://board3.com',
    backgroundSrc: 'https://board3.com/background.png',
  },
]
export const getBoards = (user_id: string) =>
  httpService<GetBoardsResponse[]>(`/api/get-boards/?user_id=${user_id}`, {
    method: 'GET',
    mock: {
      data: mockData,
      enabled: true,
    },
  })
