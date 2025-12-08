import { httpService } from '../service/http-service'

interface SaveFavoriteBoardRequest {
  boardId: string
  user_id: string
}

interface BoardInterface {
  id: string
  title: string
  isFavorite: boolean
  linkToBoard: string
  backgroundSrc: string
}
export const saveFavoriteBoard = (body: SaveFavoriteBoardRequest) =>
  httpService<BoardInterface>('/api/save-favorite-board', {
    method: 'PATCH',
    body: body,
  })
