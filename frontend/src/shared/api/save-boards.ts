import { httpService } from '../service/http-service'

interface SaveBoardRequest {
  title: string
  backgroundId: string
  user_id: string
}

interface BoardInterface {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  backgroundSrc: string
}
export const saveBoards = (body: SaveBoardRequest) =>
  httpService<BoardInterface>('/api/save-board', {
    method: 'POST',
    body: body,
  })
